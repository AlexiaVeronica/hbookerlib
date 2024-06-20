package hbookerLib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/AlexiaVeronica/hbookerLib/hbookermodel"
	"log"
	"os"
)

// SHA256 sha256 编码
func enSha256(data string) []byte {
	ret := sha256.Sum256([]byte(data))
	return ret[:]
}

func aesDecrypt(contentText string, encryptKey string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(contentText)
	if err != nil {
		return nil, err
	}
	block, ok := aes.NewCipher(enSha256(encryptKey)[:32])
	if ok != nil {
		return nil, ok
	}

	iv, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		return nil, err
	}
	blockModel, plainText := cipher.NewCBCDecrypter(block, iv), make([]byte, len(decoded))
	blockModel.CryptBlocks(plainText, decoded)
	return plainText[:(len(plainText) - int(plainText[len(plainText)-1]))], nil
}

type RSAUtil struct{}

// GetPublicKeyRefWithKey 从 base64 编码的字符串中获取公钥
func (rsaUtil *RSAUtil) GetPublicKeyRefWithKey(key string) (*rsa.PublicKey, error) {
	keyData, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 key: %v", err)
	}

	pub, err := x509.ParsePKIXPublicKey(keyData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return publicKey, nil
}

// Encrypt 使用公钥文件路径加密字符串
func (rsaUtil *RSAUtil) Encrypt(encrypt, keyFilePath string) (string, error) {
	var publicKey *rsa.PublicKey
	var err error

	if keyFilePath == "" {
		return "", errors.New("key file path is empty")
	}

	if len(encrypt) == 0 {
		return "", errors.New("encrypt string is empty")
	}

	if keyFilePath[len(keyFilePath)-4:] == ".pem" {
		publicKey, err = rsaUtil.ReadPubKeyFromPem(keyFilePath)
	} else {
		publicKey, err = rsaUtil.GetPublicKeyRefWithKey(keyFilePath)
	}

	if err != nil {
		return "", fmt.Errorf("failed to get public key: %v", err)
	}

	encryptedMessage, err := rsaUtil.EncryptString(encrypt, publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt string: %v", err)
	}

	return encryptedMessage, nil
}

// EncryptString 使用公钥加密字符串
func (rsaUtil *RSAUtil) EncryptString(str string, publicKey *rsa.PublicKey) (string, error) {
	encryptedData, err := rsaUtil.EncryptData([]byte(str), publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt data: %v", err)
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// EncryptData 使用公钥加密数据
func (rsaUtil *RSAUtil) EncryptData(data []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	keySize := publicKey.Size()
	maxChunkSize := keySize - 11 // PKCS#1 v1.5 填充

	var encryptedData []byte
	for len(data) > 0 {
		chunkSize := len(data)
		if chunkSize > maxChunkSize {
			chunkSize = maxChunkSize
		}

		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data[:chunkSize])
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt data: %v", err)
		}

		encryptedData = append(encryptedData, chunk...)
		data = data[chunkSize:]
	}

	return encryptedData, nil
}

// ReadPubKeyFromPem 从 PEM 文件中读取公钥
func (rsaUtil *RSAUtil) ReadPubKeyFromPem(filePath string) (*rsa.PublicKey, error) {
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read PEM file: %v", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return publicKey, nil
}

func (rsaUtil *RSAUtil) GetAuthenticate(authenticate *hbookermodel.Authenticate) {
	m, _ := json.Marshal(authenticate)
	rasUtil := &RSAUtil{}
	publicKey, err := rasUtil.Encrypt(string(m), publicIOSKey)
	if err != nil {
		log.Panicln("Error encrypting Authenticate:", err)
	} else {
		authenticate.SetP(publicKey)
	}
}

func SetNewVersionP(authenticate *hbookermodel.Authenticate) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	authenticate.SetRandStr(hex.EncodeToString(bytes))
	authenticate.SetSignatures(hmacKey + signaturesKey)
	h := hmac.New(sha256.New, []byte(hmacKey))
	h.Write([]byte(authenticate.GetQueryParams()))
	p := h.Sum(nil)
	authenticate.SetP(base64.StdEncoding.EncodeToString(p))

}

package hbookermodel

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Login struct {
	Tip
	Data LoginData `json:"data"`
}

type LoginData struct {
	LoginToken string     `json:"login_token"`
	UserCode   string     `json:"user_code"`
	ReaderInfo ReaderInfo `json:"reader_info"`
	PropInfo   PropInfo   `json:"prop_info"`
	IsSetYoung string     `json:"is_set_young"`
}

type Register struct {
	Tip
	Data LoginData `json:"data"`
}

type Authenticate struct {
	AppVersion  string `json:"app_version"`
	DeviceToken string `json:"device_token"`
	LoginToken  string `json:"login_token"`
	Account     string `json:"account"`
	Refresh     string `json:"refresh"`
	Signatures  string `json:"signatures"`
	RandStr     string `json:"rand_str"`
	Timestamp   string `json:"ts"`
}

func (authenticate *Authenticate) SetAppVersion(appVersion string) *Authenticate {
	// Regular expression to match semantic versioning (e.g., 1.0.0, 2.9.290)
	if regexp.MustCompile(`^\d+\.\d+\.\d+$`).MatchString(appVersion) {
		authenticate.AppVersion = appVersion
	}
	return authenticate
}
func (authenticate *Authenticate) SetDeviceToken(deviceToken string) *Authenticate {
	authenticate.DeviceToken = deviceToken
	return authenticate
}
func (authenticate *Authenticate) SetLoginToken(loginToken string) *Authenticate {
	if len(loginToken) == 32 {
		authenticate.LoginToken = loginToken
	}
	return authenticate
}
func (authenticate *Authenticate) SetAccount(account string) *Authenticate {
	if unquoted, err := strconv.Unquote(fmt.Sprintf(`"%s"`, account)); err == nil {
		account = unquoted
	}
	// Check if the (possibly decoded) string contains "书客".
	if strings.Contains(account, "书客") {
		authenticate.Account = account
	}
	return authenticate
}
func (authenticate *Authenticate) GetAuthenticate() map[string]string {
	var result map[string]string
	m, _ := json.Marshal(authenticate)
	err := json.Unmarshal(m, &result)
	if err != nil {
		log.Panicln("Error unmarshalling Authenticate:", err)
		return nil
	}
	// p
	rasUtil := &RSAUtil{}
	publicKey, err := rasUtil.Encrypt(string(m), publicIOSKey)
	if err != nil {
		log.Panicln("Error encrypting Authenticate:", err)
	} else {
		result["p"] = publicKey
	}
	return result
}
func (authenticate *Authenticate) SetRefresh(refresh string) *Authenticate {
	authenticate.Refresh = refresh
	return authenticate
}
func (authenticate *Authenticate) SetSignatures(signatures string) *Authenticate {
	authenticate.Signatures = signatures
	return authenticate
}
func (authenticate *Authenticate) SetRandStr(randStr string) *Authenticate {
	authenticate.RandStr = randStr
	return authenticate
}
func (authenticate *Authenticate) SetTimestamp(timestamp string) *Authenticate {
	authenticate.Timestamp = timestamp
	return authenticate

}

type RSAUtil struct{}

const publicIOSKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCohMLlejVLZvmFh/XFG2N5YKAjCeU08hiWUXGTUztFcUnYYhv2J1FknW/FuinK+ojveEYTNpHeXvXBjc7PXVGYLzCt+B4XW7zheehTcE8Wut3IzJd8rnIUbNpqLgqe6Ttu/X46E8wI8Xnkxlluh0wPRPIu+MmqyS1k6+2A6m/tQIDAQAB`

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

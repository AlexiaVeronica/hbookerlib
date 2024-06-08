package hbookerLib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
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

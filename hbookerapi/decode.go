package hbookerapi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// SHA256 sha256 编码
func SHA256(data []byte) []byte {
	ret := sha256.Sum256(data)
	return ret[:]
}

// AESDecrypt AES 解密
func aesDecrypt(contentText string, encryptKey []byte) ([]byte, error) {
	var iv = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if decoded, err := base64.StdEncoding.DecodeString(contentText); err == nil {
		if block, ok := aes.NewCipher(SHA256(encryptKey)[:32]); ok == nil {
			blockModel, plainText := cipher.NewCBCDecrypter(block, iv), make([]byte, len(decoded))
			blockModel.CryptBlocks(plainText, decoded)
			return plainText[:(len(plainText) - int(plainText[len(plainText)-1]))], nil
		} else {
			return nil, ok
		}
	} else {
		return nil, err
	}
}

// HbookerDecode 入口函数
func HbookerDecode(content string, encryptKey string) []byte {
	raw, ok := aesDecrypt(content, []byte(encryptKey))
	if ok == nil {
		return raw
	}
	fmt.Println("Decrypt Error, Please Check Your Key!", ok)
	return nil
}

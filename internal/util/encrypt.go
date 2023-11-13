package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	mystrings "strings"

	"github.com/gin-gonic/gin"
)

// key 16
// AES-128
// CBC
// PKCS7Padding

// AesEncrypt 加密
func AesEncrypt(orig string, key string) (string, error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 16, 24, 32
	if len(k) < 16 {
		k = append(k, bytes.Repeat([]byte{0}, 16-len(k))...)
	} else if len(k) < 24 {
		k = append(k, bytes.Repeat([]byte{0}, 24-len(k))...)
	} else if len(k) < 32 {
		k = append(k, bytes.Repeat([]byte{0}, 32-len(k))...)
	} else {
		k = k[:32]
	}
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return base64.StdEncoding.EncodeToString(cryted), nil

}

// AesDecrypt aes解密
func AesDecrypt(cryted string, key string) (string, error) {
	// 转成字节数组
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return "", err
	}
	k := []byte(key)
	// 16, 24, 32
	if len(k) < 16 {
		k = append(k, bytes.Repeat([]byte{0}, 16-len(k))...)
	} else if len(k) < 24 {
		k = append(k, bytes.Repeat([]byte{0}, 24-len(k))...)
	} else if len(k) < 32 {
		k = append(k, bytes.Repeat([]byte{0}, 32-len(k))...)
	} else {
		k = k[:32]
	}

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig), nil
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if (length - unpadding) < 0 {
		return []byte{}
	}
	return origData[:(length - unpadding)]
}

// DecryptAesEcb aes ecb 解密
func DecryptAesEcb(data, key []byte) ([]byte, error) {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	size := cip.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cip.Decrypt(decrypted[bs:be], data[bs:be])
	}
	decrypted = PKCS7UnPadding(decrypted)
	return decrypted, nil
}

// WechatGetSign 获取签名
func WechatSign(appSecret string, paramsMap gin.H) string {
	var args []string
	var keys []string
	for k := range paramsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := fmt.Sprintf("%s=%v", k, paramsMap[k])
		args = append(args, v)
	}
	baseString := mystrings.Join(args, "&")
	baseString += fmt.Sprintf("&key=%s", appSecret)
	data := []byte(baseString)
	r := md5.Sum(data)
	signedString := hex.EncodeToString(r[:])
	return mystrings.ToUpper(signedString)
}

// WechatCheckSign 检查签名
func WechatCheck(appSecret string, paramsMap gin.H) bool {
	noSignMap := gin.H{}
	for k, v := range paramsMap {
		if k != "sign" {
			noSignMap[k] = v
		}
	}
	getSign := WechatSign(appSecret, noSignMap)
	return getSign == paramsMap["sign"]
}

package main

import (
	"fmt"
	"bytes"
	"crypto/des"
	"crypto/cipher"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 3DES 加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) { 
	block, err := des.NewTripleDESCipher(key) 
	if err != nil { 
		return nil, err 
	} 
	origData = PKCS5Padding(origData, block.BlockSize()) 
	// origData = ZeroPadding(origData, block.BlockSize()) 
	blockMode := cipher.NewCBCEncrypter(block, key[:8]) 
	crypted := make([]byte, len(origData)) 
	blockMode.CryptBlocks(crypted, origData) 
	return crypted, nil 
}

// 3DES解密 
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) { 
	block, err := des.NewTripleDESCipher(key) 
	if err != nil { 
		return nil, err 
	} 
	blockMode := cipher.NewCBCDecrypter(block, key[:8]) 
	origData := make([]byte, len(crypted)) 
	// origData := crypted blockMode.CryptBlocks(origData, crypted) 
	origData = PKCS5UnPadding(origData) 
	// origData = ZeroUnPadding(origData) 
	return origData, nil 
}

func DesEncrypt(origData, key []byte) ([]byte, error) { 
	block, err := des.NewCipher(key) 
	if err != nil { 
		return nil, err 
	} 
	origData = PKCS5Padding(origData, block.BlockSize()) 
	// origData = ZeroPadding(origData, block.BlockSize()) 
	//blockMode := cipher.NewCBCEncrypter(block, key) 
	crypted := make([]byte, len(origData)) 
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以 
	// crypted := origData blockMode.CryptBlocks(crypted, origData) 
	return crypted, nil 
}

func DesDecrypt(crypted, key []byte) ([]byte, error) { 
	block, err := des.NewCipher(key) 
	if err != nil { 
		return nil, err 
	} 
	blockMode := cipher.NewCBCDecrypter(block, key) 
	origData := make([]byte, len(crypted)) 
	// origData := crypted 
	blockMode.CryptBlocks(origData, crypted) 
	origData = PKCS5Unpadding(origData) 
	// origData = ZeroUnPadding(origData) 
	return origData, nil 
}

func main() {
	key := []byte("12345678")
	origdata := []byte("88890001FFFFFFFF")
	cipher, err := DesEncrypt(origdata, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(cipher))
}

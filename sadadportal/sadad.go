package sadadportal

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func TripleDesECBEncrypt(src, key []byte) (string, error) {
	fmt.Println(key)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println("find error")
		return "", err
	}
	bs := block.BlockSize()
	origData := PKCS5Padding(src, bs)
	if len(origData)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}

	outString := base64.StdEncoding.EncodeToString(out)
	return outString, nil
}

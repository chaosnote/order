package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Base64Encode ...
func Base64Encode(plaintext string) string {
	return base64.StdEncoding.EncodeToString([]byte(plaintext))
}

// Base64Decode ...
func Base64Decode(ciphertext string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(ciphertext)
}

// MD5Encode ...
func MD5Encode(plaintext string) string {
	h := md5.New()
	io.WriteString(h, plaintext)

	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA256Encode ...
func SHA256Encode(plaintext string) string {
	h := sha256.New()
	h.Write([]byte(plaintext))
	// hex.EncodeToString(...)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HexDecode 16 進位轉為 string
func HexEncode(plaintext []byte) string {
	return hex.EncodeToString(plaintext)
}

// HexDecode 16 進位反轉為 []byte
func HexDecode(ciphertext string) ([]byte, error) {
	return hex.DecodeString(ciphertext)
}

// -----------------------------------------------

// RSAEncode
//
// warn!! return Hex
//
// call RSAInit(file_path) before using
func RSAEncode(plaintext []byte) (ciphertext string, e error) {
	var ref []byte
	ref, e = rsa.EncryptPKCS1v15(rand.Reader, &rsa_key.PublicKey, plaintext) // 加密明文信息
	if e != nil {
		return
	}
	ciphertext = fmt.Sprintf("%x", ref)
	return
}

// RSADecode
//
// call RSAInit(file_path) before using
//
// 請留意是否為 Hex 是的話需在使用 utils.HexDecode 反解
func RSADecode(ciphertext string) (plaintext []byte, e error) {
	plaintext, e = rsa.DecryptPKCS1v15(rand.Reader, rsa_key, []byte(ciphertext)) // (私)解密密文信息
	return
}

var rsa_key *rsa.PrivateKey

// RSAInit 使用前需呼叫
//
// 加密長度 {"max_limite", max/8-11}
//
// replace bool 是否取代本地檔案
func RSAInit(file_path string, max int, replace bool) {
	// 取代 且 檔案存在
	if !replace && FileExist(file_path) {
		rsa_key = read_rsa_key(file_path)
		return
	}

	var e error
	rsa_key, e = rsa.GenerateKey(rand.Reader, max)
	if e != nil {
		panic(e)
	}
	save_rsa_key(rsa_key, file_path)
}

func save_rsa_key(privateKey *rsa.PrivateKey, filename string) error {
	dir, _ := filepath.Split(filename)
	e := os.MkdirAll(dir, os.ModePerm)
	if e != nil {
		return e
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey) // 格式化私鑰
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	file, e := os.Create(filename) // 寫入文件
	if e != nil {
		return e
	}
	defer file.Close()

	e = pem.Encode(file, pemBlock)
	if e != nil {
		return e
	}
	return nil
}

func read_rsa_key(filename string) *rsa.PrivateKey {
	privateKeyPEM, e := os.ReadFile(filename) // 讀取 pem 檔案
	if e != nil {
		panic(e)
	}
	block, _ := pem.Decode(privateKeyPEM)                   // pem 轉私鑰解碼
	privateKey, e := x509.ParsePKCS1PrivateKey(block.Bytes) // 解碼數據轉為 rsa 私鑰
	if e != nil {
		panic(e)
	}
	return privateKey
}

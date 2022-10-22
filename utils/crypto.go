package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

var ENCRYPT_SEED = uint64(11468049314633205968)
var DISPATCH_KEY, DISPATCH_SEED, ENCRYPT_KEY, ENCRYPT_SEED_BUFFER []byte
var CUR_SIGNING_KEY, AUTH_PRIVATE_KEY *rsa.PrivateKey
var CUR_OS_ENCRYPT_KEY, CUR_CN_ENCRYPT_KEY *rsa.PublicKey

func init() {
	key, err := ioutil.ReadFile("./data/keys/DispatchKey.bin")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! DispatchKey.bin 文件缺失, error: ", err.Error())
		return
	}
	DISPATCH_KEY = key

	key, err = ioutil.ReadFile("./data/keys/DispatchSeed.bin")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! DispatchSeed.bin 文件缺失, error: ", err.Error())
		return
	}
	DISPATCH_SEED = key

	key, err = ioutil.ReadFile("./data/keys/SecretKey.bin")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! SecretKey.bin 文件缺失, error: ", err.Error())
		return
	}
	ENCRYPT_KEY = key

	key, err = ioutil.ReadFile("./data/keys/SecretKeyBuffer.bin")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! SecretKeyBuffer.bin 文件缺失, error: ", err.Error())
		return
	}
	ENCRYPT_SEED_BUFFER = key

	file, err := ioutil.ReadFile("./data/keys/SigningKey.pem")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! SigningKey.pem 文件缺失, error: ", err.Error())
		return
	}
	block, _ := pem.Decode(file)
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Panicln("[Crypto] 解析失败! SigningKey.pem 无法解析, error: ", err.Error())
		return
	}
	CUR_SIGNING_KEY = privateKey.(*rsa.PrivateKey)

	file, err = ioutil.ReadFile("./data/keys/AuthPrivateKey.pem")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! AuthPrivateKey.pem 文件缺失, error: ", err.Error())
		return
	}
	block, _ = pem.Decode(file)
	privateKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Panicln("[Crypto] 解析失败! AuthPrivateKey.pem 无法解析, error: ", err.Error())
		return
	}
	AUTH_PRIVATE_KEY = privateKey.(*rsa.PrivateKey)

	file, err = ioutil.ReadFile("./data/keys/OSCB.pem")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! OSCB.pem 文件缺失, error: ", err.Error())
		return
	}
	block, _ = pem.Decode(file)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Panicln("[Crypto] 解析失败! OSCB.pem 无法解析, error: ", err.Error())
		return
	}
	CUR_OS_ENCRYPT_KEY = publicKey.(*rsa.PublicKey)

	file, err = ioutil.ReadFile("./data/keys/OSCN.pem")
	if err != nil {
		log.Panicln("[Crypto] 读取失败! OSCN.pem 文件缺失, error: ", err.Error())
		return
	}
	block, _ = pem.Decode(file)
	publicKey, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Panicln("[Crypto] 解析失败! OSCN.pem 无法解析, error: ", err.Error())
		return
	}
	CUR_CN_ENCRYPT_KEY = publicKey.(*rsa.PublicKey)
}

// XOR加密
func XOR(data []byte, key []byte) []byte {
	for i := range data {
		data[i] ^= key[i%len(key)]
	}
	return data
}

// 公钥加密-分段
func RsaEncryptBlock(src []byte, publicKey *rsa.PublicKey) (bytesEncrypt []byte, err error) {
	keySize, srcSize := publicKey.Size(), len(src)
	//logs.Debug("密钥长度：", keySize, "\t明文长度：\t", srcSize)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt = buffer.Bytes()
	return
}

// 私钥解密-分段
func RsaDecryptBlock(src []byte, privateKey *rsa.PrivateKey) (bytesDecrypt []byte, err error) {
	keySize, srcSize := privateKey.Size(), len(src)
	//logs.Debug("密钥长度：", keySize, "\t密文长度：\t", srcSize)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt = buffer.Bytes()
	return
}

func AbilityHash(str string) int {
	num1 := 0
	num2 := 0
	for num2 < len(str) {
		num1 = int(str[num2]) + 131*num1
		num2++
	}
	return num1
}

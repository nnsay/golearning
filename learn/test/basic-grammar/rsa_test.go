package golearning_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

func RSAGenKey(bits int) error {
	/*
		生成私钥
	*/
	//1、使用RSA中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	//2、通过X509标准将得到的RAS私钥序列化为：ASN.1 的DER编码字符串
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	//3、将私钥字符串设置到pem格式块中
	block1 := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateStream,
	}
	//4、通过pem将设置的数据进行编码，并写入磁盘文件
	fPrivate, err := os.Create("privateKey.pem")
	if err != nil {
		return err
	}
	defer fPrivate.Close()
	err = pem.Encode(fPrivate, &block1)
	if err != nil {
		return err
	}

	/*
		生成公钥
	*/
	publicKey := privateKey.PublicKey
	publicStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	//publicStream:=x509.MarshalPKCS1PublicKey(&publicKey)
	block2 := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicStream,
	}
	fPublic, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	defer fPublic.Close()
	pem.Encode(fPublic, &block2)
	return nil
}
func TestGenerateKey(t *testing.T) {
	RSAGenKey(1024)
}

func TestRSA(t *testing.T) {
	RSAPublicBytes := RSAReadKeyFromFile("/home/nguser/github/ngrepos/ngiq-sso-deploy/asset/go-ldap-admin-server-pub.pem")
	RSAPrivateBytes := RSAReadKeyFromFile("/home/nguser/github/ngrepos/ngiq-sso-deploy/asset/go-ldap-admin-server-priv.pem")
	// RSAPublicBytes := RSAReadKeyFromFile("/Users/wangjian/Workspace/practice/golearning/learn/test/basic-grammar/publicKey.pem")
	// RSAPrivateBytes := RSAReadKeyFromFile("/Users/wangjian/Workspace/practice/golearning/learn/test/basic-grammar/privateKey.pem")
	pass, _ := RSAEncrypt([]byte("123456"), RSAPublicBytes)
	hashPass := string(pass)
	plainPass, err := RSADecrypt([]byte(hashPass), RSAPrivateBytes)
	t.Logf("%#v", string(plainPass))
	t.Logf("%#v", err)
}
func RSAReadKeyFromFile(filename string) []byte {
	f, err := os.Open(filename)
	var b []byte

	if err != nil {
		return b
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b = make([]byte, fileInfo.Size())
	_, err = f.Read(b)
	if err != nil {
		return b
	}
	return b
}

// RSA加密
func RSAEncrypt(data, publicBytes []byte) ([]byte, error) {
	var res []byte
	// 解析公钥
	block, _ := pem.Decode(publicBytes)

	if block == nil {
		return res, fmt.Errorf("无法加密, 公钥可能不正确")
	}

	// 使用X509将解码之后的数据 解析出来
	// x509.MarshalPKCS1PublicKey(block):解析之后无法用，所以采用以下方法：ParsePKIXPublicKey
	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return res, fmt.Errorf("无法加密, 公钥可能不正确, %v", err)
	}
	// 使用公钥加密数据
	pubKey := keyInit.(*rsa.PublicKey)
	res, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		return res, fmt.Errorf("无法加密, 公钥可能不正确, %v", err)
	}
	// 将数据加密为base64格式
	return []byte(EncodeStr2Base64(string(res))), nil
}

// 对数据进行解密操作
func RSADecrypt(base64Data, privateBytes []byte) ([]byte, error) {
	var res []byte
	// 将base64数据解析
	data := []byte(DecodeStrFromBase64(string(base64Data)))
	// 解析私钥
	block, _ := pem.Decode(privateBytes)
	if block == nil {
		return res, fmt.Errorf("无法解密, 私钥可能不正确,解析私钥失败")
	}
	// 还原数据
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return res, fmt.Errorf("无法解密, 私钥可能不正确,解析PKCS失败 %v", err)
	}
	res, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if err != nil {
		return res, fmt.Errorf("无法解密, 私钥可能不正确,解密PKCS1v15失败 %v", err)
	}
	return res, nil
}

// 加密base64字符串
func EncodeStr2Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解密base64字符串
func DecodeStrFromBase64(str string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(decodeBytes)
}

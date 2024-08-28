package key

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"myServer/global"
)

// LoadPublicKey 加载PEM公钥并解码
func LoadPublicKey() (*rsa.PublicKey, error) {
	path := global.Config.Jwt.PublicKeyPath
	pemData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read public key file: %v", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("expected RSA PUBLIC KEY but got %s", block.Type)
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse RSA public key: %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not a valid RSA public key")
	}

	return rsaPublicKey, nil
}

func LoadPrivateKey() (*rsa.PrivateKey, error) {
	path := global.Config.Jwt.PrivateKeyPath // 从全局配置中获取公钥存储路径

	// 读取私钥文件
	pemData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read public key file: %v", err)
	}

	// 解码PEM格式数据
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// 解析公钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse RSA private key: %v", err)
	}

	return privateKey, nil
}

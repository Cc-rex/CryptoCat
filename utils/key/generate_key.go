package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateKeyPair generates a new key pair if they do not exist
func GenerateKeyPair(privateKeyPath, publicKeyPath string) error {
	// Check if the private key already exists
	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
		// Generate RSA private key
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return err
		}

		// Save private key
		privateFile, err := os.Create(privateKeyPath)
		if err != nil {
			return err
		}
		defer privateFile.Close()

		privateKeyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}
		if err := pem.Encode(privateFile, privateKeyPEM); err != nil {
			return err
		}

		// Generate and save public key
		publicKey := &privateKey.PublicKey
		publicFile, err := os.Create(publicKeyPath)
		if err != nil {
			return err
		}
		defer publicFile.Close()

		publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return err
		}

		publicKeyPEM := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		}
		if err := pem.Encode(publicFile, publicKeyPEM); err != nil {
			return err
		}
		fmt.Println("生成密钥成功")
	} else {
		fmt.Println("密钥已存在，无需重新生成")
	}
	return nil
}

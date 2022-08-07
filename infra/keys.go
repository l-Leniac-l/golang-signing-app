package infra

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func LoadPublicKey() (*ecdsa.PublicKey, error) {
	pubKey := os.Getenv("PUBLIC_KEY")
	block, _ := pem.Decode([]byte(pubKey))

	if block == nil {
		return nil, errors.New("failed to load public key")
	}

	genericPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	pk := genericPublicKey.(*ecdsa.PublicKey)

	return pk, nil
}

func LoadPrivateKey() (*ecdsa.PrivateKey, error) {
	privKey := os.Getenv("PRIVATE_KEY")
	block, _ := pem.Decode([]byte(privKey))

	if block == nil {
		return nil, errors.New("failed to load private key")
	}

	pk, err := x509.ParseECPrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return pk, nil
}

package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

type SignIdentity struct {
	user         *User
	documentHash []byte
	signature    string
}

func (s *SignIdentity) SetSignature(signature string) {
	s.signature = signature
}

func (s *SignIdentity) Sign() (string, error) {
	pk, err := LoadPrivateKey()

	if err != nil {
		return "", err
	}

	h := sha512.New()

	h.Write([]byte(s.user.ToString()))
	h.Write(s.documentHash)

	hash := h.Sum(nil)

	sign, err := ecdsa.SignASN1(rand.Reader, pk, hash)

	if err != nil {
		return "", err
	}

	bSign := base64.StdEncoding.EncodeToString(sign)

	s.signature = bSign

	return bSign, nil
}

func (s *SignIdentity) Validate() (bool, error) {
	pk, err := LoadPublicKey()

	if err != nil {
		return false, err
	}

	h := sha512.New()

	h.Write([]byte(s.user.ToString()))
	h.Write(s.documentHash)

	hash := h.Sum(nil)

	bSign, err := base64.StdEncoding.DecodeString(s.signature)

	if err != nil {
		return false, err
	}

	return ecdsa.VerifyASN1(pk, hash, bSign), nil
}

func NewSignIdentity(user *User, documentHash []byte) *SignIdentity {
	return &SignIdentity{
		user:         user,
		documentHash: documentHash,
	}
}

package services

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"

	"github.com/l-leniac-l/golang-signing-app/domain/entities"
	"github.com/l-leniac-l/golang-signing-app/infra"
)

type SignIdentity struct {
	Identity     *entities.Identity `json:"identity"`
	DocumentHash []byte             `json:"documentHash"`
	Signature    string             `json:"signature"`
}

func (s *SignIdentity) SetSignature(signature string) {
	s.Signature = signature
}

func (s *SignIdentity) Sign() (string, error) {
	pk, err := infra.LoadPrivateKey()

	if err != nil {
		return "", err
	}

	h := sha512.New()

	h.Write([]byte(s.Identity.ToString()))
	h.Write(s.DocumentHash)

	hash := h.Sum(nil)

	sign, err := ecdsa.SignASN1(rand.Reader, pk, hash)

	if err != nil {
		return "", err
	}

	bSign := base64.StdEncoding.EncodeToString(sign)

	s.Signature = bSign

	return bSign, nil
}

func (s *SignIdentity) Validate() (bool, error) {
	pk, err := infra.LoadPublicKey()

	if err != nil {
		return false, err
	}

	h := sha512.New()

	h.Write([]byte(s.Identity.ToString()))
	h.Write(s.DocumentHash)

	hash := h.Sum(nil)

	bSign, err := base64.StdEncoding.DecodeString(s.Signature)

	if err != nil {
		return false, err
	}

	return ecdsa.VerifyASN1(pk, hash, bSign), nil
}

func NewSignIdentity(identity *entities.Identity, documentHash []byte) *SignIdentity {
	return &SignIdentity{
		Identity:     identity,
		DocumentHash: documentHash,
	}
}

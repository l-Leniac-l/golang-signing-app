package services

import (
	"crypto/sha512"
	"testing"

	"github.com/joho/godotenv"
	"github.com/l-leniac-l/golang-signing-app/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestSignNoPrivateKey(t *testing.T) {
	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	identity := &entities.Identity{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(identity, hash)

	_, err := si.Sign()

	got := err.Error()

	want := "failed to load private key"

	assert.Equal(t, want, got)
}

func TestSignSuccess(t *testing.T) {
	envMap, _ := godotenv.Read("../../.env")

	for k, v := range envMap {
		t.Setenv(k, v)
	}
	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	identity := &entities.Identity{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(identity, hash)

	sign, err := si.Sign()

	if assert.NoError(t, err) {
		assert.Equal(t, sign, si.Signature)
	}
}

func TestSetSignature(t *testing.T) {
	si := SignIdentity{}

	si.SetSignature("MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTkTZI0=")

	got := si.Signature
	want := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTkTZI0="

	assert.Equal(t, want, got)
}

func TestValidateNoPublicKey(t *testing.T) {
	si := SignIdentity{}

	sign := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTkTZI0="

	si.SetSignature(sign)

	_, err := si.Validate()

	got := err.Error()

	want := "failed to load public key"

	assert.Equal(t, want, got)
}

func TestValidateBadSignature(t *testing.T) {
	envMap, _ := godotenv.Read("../../.env")

	for k, v := range envMap {
		t.Setenv(k, v)
	}

	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	identity := &entities.Identity{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(identity, hash)

	sign := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTZI0="

	si.SetSignature(sign)

	_, err := si.Validate()

	got := err.Error()

	want := "illegal base64"

	assert.Contains(t, got, want)
}
func TestValidateSuccess(t *testing.T) {
	envMap, _ := godotenv.Read("../../.env")

	for k, v := range envMap {
		t.Setenv(k, v)
	}
	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	identity := &entities.Identity{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(identity, hash)

	sign := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTkTZI0="

	si.SetSignature(sign)

	valid, _ := si.Validate()

	assert.Equal(t, true, valid)
}

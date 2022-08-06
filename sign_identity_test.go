package main

import (
	"crypto/sha512"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestSignNoPrivateKey(t *testing.T) {
	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	user := &User{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(user, hash)

	_, err := si.Sign()

	got := err.Error()

	want := "failed to load private key"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSign(t *testing.T) {
	envMap, _ := godotenv.Read()

	for k, v := range envMap {
		t.Setenv(k, v)
	}
	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	user := &User{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(user, hash)

	sign, err := si.Sign()

	if err != nil {
		t.Errorf("error while signing document: %s", err)
	}

	if sign != si.signature {
		t.Errorf("got %s, wanted %s", sign, si.signature)
	}
}

func TestSetSignature(t *testing.T) {
	si := SignIdentity{}

	si.SetSignature("test")

	got := si.signature
	want := "test"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestValidateNoPublicKey(t *testing.T) {
	si := SignIdentity{}

	sign := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTkTZI0="

	si.SetSignature(sign)

	_, err := si.Validate()

	got := err.Error()

	want := "failed to load public key"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestValidateBadSignature(t *testing.T) {
	envMap, _ := godotenv.Read()

	for k, v := range envMap {
		t.Setenv(k, v)
	}

	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	user := &User{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(user, hash)

	sign := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTZI0="

	si.SetSignature(sign)

	_, err := si.Validate()

	got := err.Error()

	want := "illegal base64"

	if !strings.Contains(got, want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestValidate(t *testing.T) {
	envMap, _ := godotenv.Read()

	for k, v := range envMap {
		t.Setenv(k, v)
	}
	// Creating a document hash
	h := sha512.New()
	h.Write([]byte("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"))
	hash := h.Sum(nil)

	user := &User{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	si := NewSignIdentity(user, hash)

	sign := "MEUCIQC+iyknVMK3L58jjTikEotji7n5bXVTHMTQpDxBNqb4BwIgH3btjoj4dCG6EKRS4zpeUDMF3GbFiqo9vVutVTkTZI0="

	si.SetSignature(sign)

	valid, _ := si.Validate()

	if valid != true {
		t.Errorf("got %v, wanted %v", valid, true)
	}
}

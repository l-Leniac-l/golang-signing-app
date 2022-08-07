package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSignRouteInvalidBody(t *testing.T) {
	router := SetupRouter()

	reqBody := ""

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/identities/sign", bytes.NewReader(body))

	router.ServeHTTP(w, req)

	var m map[string]interface{}

	json.NewDecoder(w.Body).Decode(&m)

	expected := map[string]interface{}{
		"message": "invalid request body",
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, m)
}

func TestSignRouteIdentityNotFound(t *testing.T) {
	router := SetupRouter()

	reqBody := map[string]interface{}{
		"identityId":   "f54476b7-440a-41f3-b271-5d47aeecff35",
		"documentHash": "5a9ed5765f61ba319f6a3eea81469a3a099b7433a87c96da0989b713a616c5ab67a3067f231bc0f99097ade53e0ac47d8e116436a734a18b074a626c3fee7afc",
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/identities/sign", bytes.NewReader(body))

	router.ServeHTTP(w, req)

	var m map[string]interface{}

	json.NewDecoder(w.Body).Decode(&m)

	expected := map[string]interface{}{
		"message": fmt.Sprintf(
			"cannot found identity with id: %s",
			reqBody["identityId"],
		),
	}

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expected, m)
}

func TestSignIdentityInvalidPrivateKey(t *testing.T) {
	t.Setenv("PRIVATE_KEY", "invalid key")

	router := SetupRouter()

	reqBody := map[string]interface{}{
		"identityId":   "f0b48033-dbb2-4bd5-b24d-8f8763e9461f",
		"documentHash": "5a9ed5765f61ba319f6a3eea81469a3a099b7433a87c96da0989b713a616c5ab67a3067f231bc0f99097ade53e0ac47d8e116436a734a18b074a626c3fee7afc",
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/identities/sign", bytes.NewReader(body))

	router.ServeHTTP(w, req)

	var m map[string]interface{}

	json.NewDecoder(w.Body).Decode(&m)

	expected := map[string]interface{}{
		"message": "failed to sign document",
		"error":   "failed to load private key",
	}

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, m)
}

func TestSignIdentitySuccess(t *testing.T) {
	envMap, _ := godotenv.Read("../../.env")

	for k, v := range envMap {
		t.Setenv(k, v)
	}

	router := SetupRouter()

	reqBody := map[string]interface{}{
		"identityId":   "f0b48033-dbb2-4bd5-b24d-8f8763e9461f",
		"documentHash": "5a9ed5765f61ba319f6a3eea81469a3a099b7433a87c96da0989b713a616c5ab67a3067f231bc0f99097ade53e0ac47d8e116436a734a18b074a626c3fee7afc",
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/identities/sign", bytes.NewReader(body))

	router.ServeHTTP(w, req)

	var m map[string]interface{}

	json.NewDecoder(w.Body).Decode(&m)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, m, "signedIdentity")
	assert.Contains(t, m["signedIdentity"], "signature")
}

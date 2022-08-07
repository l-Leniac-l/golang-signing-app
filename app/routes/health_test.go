package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)

	router.ServeHTTP(w, req)

	var m map[string]interface{}

	json.NewDecoder(w.Body).Decode(&m)

	expected := map[string]interface{}{
		"status": "running",
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, m)
}

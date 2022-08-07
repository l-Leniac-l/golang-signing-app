package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityToString(t *testing.T) {
	identity := &Identity{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	got := identity.ToString()
	want := "John,Doe,1995-03-15"

	assert.Equal(t, want, got)
}

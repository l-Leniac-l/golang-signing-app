package repositories

import (
	"testing"

	"github.com/l-leniac-l/golang-signing-app/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestByIdNotFound(t *testing.T) {
	ir := NewIdentityRepository()

	got := ir.GetById("69e4a5ff-900b-447e-8978-206e3fc33447")

	var want *entities.Identity = nil

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
func TestGetByIdSuccess(t *testing.T) {
	want := &entities.Identity{
		Id:          "f0b48033-dbb2-4bd5-b24d-8f8763e9461f",
		FirstName:   "Doe",
		LastName:    "John",
		DateOfBirth: "2001-05-03",
	}

	ir := NewIdentityRepository()

	got := ir.GetById("f0b48033-dbb2-4bd5-b24d-8f8763e9461f")

	assert.Equal(t, want, got)
}

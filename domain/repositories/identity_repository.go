package repositories

import "github.com/l-leniac-l/golang-signing-app/domain/entities"

type IdentityRepository struct {
	identities *[]entities.Identity
}

func (ir *IdentityRepository) GetById(id string) *entities.Identity {
	for _, item := range *ir.identities {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func NewIdentityRepository() *IdentityRepository {
	identities := &[]entities.Identity{
		{
			Id:          "e7616ccb-e794-472b-8111-ea2a05794a96",
			FirstName:   "John",
			LastName:    "Doe",
			DateOfBirth: "1995-03-15",
		},
		{
			Id:          "f0b48033-dbb2-4bd5-b24d-8f8763e9461f",
			FirstName:   "Doe",
			LastName:    "John",
			DateOfBirth: "2001-05-03",
		},
		{
			Id:          "d264d4a0-2679-48c0-95ec-961caa81d277",
			FirstName:   "Daven",
			LastName:    "Alex",
			DateOfBirth: "1993-08-22",
		},
		{
			Id:          "91b00555-4c86-42df-b199-4a954f5310b2",
			FirstName:   "Alex",
			LastName:    "Daven",
			DateOfBirth: "1957-12-11",
		},
	}

	return &IdentityRepository{
		identities: identities,
	}
}

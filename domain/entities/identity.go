package entities

import "fmt"

type Identity struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (u *Identity) ToString() string {
	return fmt.Sprintf("%s,%s,%s", u.FirstName, u.LastName, u.DateOfBirth)
}

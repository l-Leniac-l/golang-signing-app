package main

import "fmt"

type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("%s,%s,%s", u.FirstName, u.LastName, u.DateOfBirth)
}

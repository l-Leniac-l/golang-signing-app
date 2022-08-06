package main

import "testing"

func TestUserToString(t *testing.T) {
	user := &User{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1995-03-15",
	}

	got := user.ToString()
	want := "John,Doe,1995-03-15"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

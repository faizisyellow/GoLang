package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Method (function inside struct)
func (user User) OutputUserDetails() {

	// in struct to direferencing does automaticly  by go
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

// pasing by referance
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

// Method creation / constructor. (blue print creating)
func New(firstname, lastname, birthdate string) (*User, error) {
	// can do validation too.
	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("required")
	}

	// creating struct
	return &User{
		firstName: firstname,
		lastName:  lastname,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

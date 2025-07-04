package structs

import (
	"errors"
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	age       int
}

type Admin struct {
	email    string
	password string
	user     User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		user: User{
			firstName: "admin",
			lastName:  "admin",
			age:       32,
		},
	}
}

func NewUser(firstName string, lastName string, age int) (*User, error) {
	if firstName == "" || lastName == "" || age <= 0 {
		return nil, errors.New("missing required fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}, nil
}

func (userDetails User) ShowUserDetails() {
	fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.age)
}

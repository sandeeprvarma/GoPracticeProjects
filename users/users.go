package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var firstName, lastName string
	var age int
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&age)

	var userDetails = user{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}

	showUserDetails(&userDetails)
}

func showUserDetails(userDetails *user) {
	fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.age)
}

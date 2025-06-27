package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	age       int
}

func newUser(firstName string, lastName string, age int) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}
func main() {
	var firstName, lastName string
	var age int
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&age)

	var userDetails = newUser(firstName, lastName, age)

	userDetails.showUserDetails()
	// showUserDetails(&userDetails)
}

//	func showUserDetails(userDetails *user) {
//		fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.age)
//	}
func (userDetails user) showUserDetails() {
	fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.age)
}

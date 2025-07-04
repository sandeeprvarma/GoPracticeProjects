package main

import (
	"fmt"
	"users/structs"
)

func main() {
	var firstName, lastName string
	var age int
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&age)

	var userDetails *structs.User
	userDetails, err := structs.NewUser(firstName, lastName, age)
	adminDetail := structs.NewAdmin("a@b.c", "abcd")
	fmt.Printf("adminDetail: %v\n", adminDetail)

	if err != nil {
		fmt.Println(err)
		return
	}
	userDetails.ShowUserDetails()
	// showUserDetails(&userDetails)
}

//	func showUserDetails(userDetails *user) {
//		fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.age)
//	}

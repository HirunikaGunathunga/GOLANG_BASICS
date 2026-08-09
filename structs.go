package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func structFunc() {

	userFirstName := getUserData("First name here: ")
	userLastName := getUserData("Last name here: ")
	userBirthDate := getUserData("Birth day here: ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	} //can create a NULL Struct value
	outputUserData(appUser)
}

func outputUserData(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func getUserData(names string) string {

	fmt.Print(names)
	var data string
	fmt.Scan(&data)
	return data
}

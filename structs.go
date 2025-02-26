package main

import (
	"fmt"

	"example.com/first-app/packages/communication"
	"example.com/first-app/packages/user"
)

func structs() {

	userFirstName := communication.TakeInput("Please enter your first name:")
	userLastName := communication.TakeInput("Please enter your last name:")
	userBirthDate := communication.TakeInput("Please enter your birthdate:")
	var userStruct *user.UserType
	userStruct, userErr := user.NewUserConstructor(userFirstName, userLastName, userBirthDate)
	if userErr != nil {
		panic(userErr)
	} else {
		fmt.Println("User created successfully!")
		userStruct.PrintValues()
	}
	var adminStruct *user.AdminType
	adminStruct, adminErr := user.NewAdminConstrucor("dummyEmaail@email.com", "password", userStruct)
	if adminErr != nil {
		panic(adminErr)
	} else {
		fmt.Println("Admin created successfully!")
		adminStruct.PrintAdmin(userStruct)
	}
}

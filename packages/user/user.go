package user

import (
	"errors"
	"fmt"
)

type UserType struct {
	firstName string
	lastName  string
	birthday  string
}

type AdminType struct {
	email    string
	password string
	user     UserType // inheritance in GO, called embading
}

func (user *UserType) PrintValues() { // method
	fmt.Printf("User's First Name is ---> %s\n User's Last name is ---> %s\n User's Birthdate is ---> %s\n",
		user.firstName, user.lastName, user.birthday)
}

func (admin *AdminType) PrintAdmin(user *UserType) {
	fmt.Printf("Admin's Email is ---> %s\n Admin's Password is ---> %s\n Admin's firstName is ---> %s\n Admin's lastName is --->%s/n Admin's Birthdate is ---> %s\n",
		admin.email, admin.password, admin.user.firstName, admin.user.lastName, admin.user.birthday)
}

func (user *UserType) ClearValues() { //passing pointer varaible to ensure original value is mutated not the copied 😉
	user.firstName = ""
	user.lastName = ""
}

func NewAdminConstrucor(email string, password string, user *UserType) (*AdminType, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and password are required")
	}
	return &AdminType{
		email:    email,
		password: password,
		user: UserType{
			firstName: user.firstName, //either take this input or hard code them
			lastName:  user.lastName,
			birthday:  user.birthday,
		},
	}, nil
}

func NewUserConstructor(firstName string, lastName string, birthdate string) (*UserType, error) { //constructor function that helps to create structs seemlesly
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First and last names and birth dates are required")
	}
	return &UserType{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthdate,
	}, nil
}

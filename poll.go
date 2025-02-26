package main

import (
	"fmt"

	"example.com/first-app/packages/voting"
)

func poll() {
	fmt.Println("Please enter your age")
	var age string
	fmt.Scan(&age)
	var canVote bool = voting.CanPersonVote(&age)

	if canVote {
		fmt.Println("Please proceed to vote")
	} else {
		fmt.Println("Sorry, you are not allowed to vote")
	}
}

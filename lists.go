package main

import "fmt"

type Product struct {
	title string
	id    string
	price string
}

func list() {
	hobbies := [3]string{"Exploring New place", "Exploring new food", "Gaming"}
	fmt.Printf("1st Element: %s\n2nd and 3rd Element: %s\n", hobbies[0], hobbies[1:2])
	fmt.Printf("First and Second Element: %s\n", hobbies[0:2])

	goals := []string{}
	goals = append(goals, "Spend the 5 lakh bonus wisely", "Save 30lakh for my wedding")
	goals[1] = "Spend the 5 lakh bonus very wisely and buy goods"

	products := []Product{}

	products = append(products, Product{
		title: "Book",
		id:    "12345",
		price: "200",
	}, Product{
		title: "Mobile",
		id:    "67890",
		price: "5000",
	}, Product{
		title: "Laptop",
		id:    "09876",
		price: "10000",
	})

	fmt.Printf("Products: %+v\nHobbies:%s\nGoals: %s\n", products, hobbies, goals)

}

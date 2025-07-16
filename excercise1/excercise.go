package main

import (
	"fmt"
)

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	hobbies := [3]string{
		"running",
		"sleeping",
		"listening music",
	}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	newSlicedHobbies := hobbies[1:]
	fmt.Println(hobbies[:2])
	fmt.Println(newSlicedHobbies)
	// newSlicedHobbies = newSlicedHobbies[1:]
	goals := []string{
		"learn go",
		"build min 5 to 10 projects",
	}

	fmt.Println(goals)
	goals = append(goals, "build a microservice")
	fmt.Println(goals)

	var newProduct = []Product{}
	newProduct = append(newProduct, Product{1, "title1", 1.1}, Product{1, "title2", 1.1})
	fmt.Println(newProduct)
}

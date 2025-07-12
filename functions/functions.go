package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	myFn := trasformNumber(2)
	newNum := calculateNumber(&number, myFn)
	fmt.Println(newNum)
}

func calculateNumber(number *int, calculateFun func(int) int) int {
	return calculateFun((*number))
}

func trasformNumber(number int) func(int) int {
	if number == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

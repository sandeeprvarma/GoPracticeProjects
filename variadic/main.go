package main

import "fmt"

func main() {
	total := calculateSum(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}

func calculateSum(number ...int) int {
	sum := 0
	for _, val := range number {
		sum += val
	}
	return sum
}

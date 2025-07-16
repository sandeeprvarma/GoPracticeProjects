package main

import "fmt"

func main() {
	var productNames = [3]string{"abc", "xyz", "lmn"}
	prices := [4]float64{1.0, 2.1, 3.2, 1}
	fmt.Println(productNames, prices, prices[2], prices[1:3])
}

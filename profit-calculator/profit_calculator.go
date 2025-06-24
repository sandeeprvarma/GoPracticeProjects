package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Scan(&revenue)
	fmt.Scan(&expense)
	fmt.Scan(&taxRate)
	//earnings before taxes
	ebt := revenue - expense
	profit := ebt - (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ratio)

	// fmt.Printf("Revenue Value: %.1f \n Expense value: %.2f", revenue, expense)
	// fmt.Printf(`Revenue Value: %.1f \n
	// Expense value: %.2f`, revenue, expense)
}

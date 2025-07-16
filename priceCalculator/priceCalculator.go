package main

import "fmt"

func main() {
	prices := []float64{1.2, 44.5, 88.0}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	var results map[float64][]float64 = make(map[float64][]float64, len(prices))
	for _, taxRate := range taxRates {
		withTaxPrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			withTaxPrice[priceIndex] = price * (1 + taxRate)
		}
		results[taxRate] = withTaxPrice
	}
	fmt.Println(results)
}

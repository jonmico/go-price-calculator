package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))

		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (taxRate + 1)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}

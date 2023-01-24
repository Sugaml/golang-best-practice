package main

import "fmt"

//Not properly checking if a slice is empty
func main() {
	handlePrices("1")
}

func handlePrices(name string) {
	amounts := getAmounts(name)
	//Checks if the amounts slice is nil
	if len(amounts) != 0 {
		handle(amounts)
	}
}

func getAmounts(name string) []float32 {
	//Initializes the amounts slice
	amounts := make([]float32, 0)
	//Returns amounts if the provided name is empty
	if name == "" {
		return nil
	}
	// Add elements to amounts
	return amounts
}

func handle(amounts []float32) {
	//Do some operations
	fmt.Println(amounts)
}

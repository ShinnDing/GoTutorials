// Go's map data structure
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_12/maps.go
package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70, // Must have trailing comma in multi line
	}

	// Number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Get zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	// Use two values to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TLSA not found")
	} else {
		fmt.Println(value)
	}

	// Set
	stocks["TSLA"] = 822.12
	fmt.Println(stocks)

	// Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}

}

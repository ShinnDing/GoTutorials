// Strings formatting
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_14/words.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	fmt.Println(text)

	words := strings.Fields(text)
	counts := map[string]int{} // word -> count
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)

}

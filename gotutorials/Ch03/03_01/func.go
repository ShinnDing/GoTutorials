// Functions
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch03/03_01/func.go
package main

import (
	"fmt"
)

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)

}

// adds a to b
func add(a int, b int) int {
	return a + b
}

// divmod returs quotient and remainder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

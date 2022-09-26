// Fizz Buzz Challenge solution
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_05/fizzbuzz.go
package main

import (
	"fmt"
)

func main() {

	// for every number from 1 to 20
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			// if the number is divisible by 3 and 5 print fizz buzz
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			// else if the number is divisible by 3 print fizz
			fmt.Println("fizz")
		} else if i%5 == 0 {
			// else if the number is divisble by 5 print buzz
			fmt.Println("buzz")
		} else {
			// else print the number
			fmt.Println(i)
		}
	}

}

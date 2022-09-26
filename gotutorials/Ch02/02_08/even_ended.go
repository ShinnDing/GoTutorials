/*
Solution for Even-ended numbers Challenge
An "even ended number" is a number who's first and last digits are the same.
Count how many "even ended numbers" there are with a multiplication of 4 digit numbers.
*/
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_08/even_ended.go
package main

import (
	"fmt"
)

func main() {
	count := 0

	// for every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { // don't count twice
			n := a * b

			// if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// count = count + 1
				count++
			}
		}
	}

	// print count
	fmt.Println(count)

}

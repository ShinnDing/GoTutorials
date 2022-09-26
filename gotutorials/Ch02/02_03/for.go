// "for" loop examples
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_03/for.go
package main

import (
	"fmt"
)

func main() {
	// traditional for-loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("____")
	// break for-loop in the middle
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("____")
	// continue (go to next loop without execuding code)
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	// equivalent to while-loop
	fmt.Println("____")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	// equivalent to while-true or infinite loop
	fmt.Println("____")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}

}

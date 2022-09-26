// Go slices
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_09/slices.go
package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons)) // 3

	fmt.Println("____")
	// 0 indexing
	fmt.Println(loons[1]) // daffy

	fmt.Println("____")
	// slices
	fmt.Println(loons[1:]) // [daffy taz]

	fmt.Println("____")
	// for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("____")
	// Single value range
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("____")
	// Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("____")
	// Double value range, ignore index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("____")
	// append
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}

// Calculate maximal value in a slice
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_11/max.go
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	// fmt.Println(nums)

	max := nums[0] // Initialize max with first value
	// [1:] skips the first value
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}

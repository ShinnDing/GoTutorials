// select example
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch06/06_05/select.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	fmt.Println("____")

	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout")
	}
}

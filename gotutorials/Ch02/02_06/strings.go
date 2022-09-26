// string examples
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch02/02_06/strings.go
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// Strings in go are immutable
	// this results in an error: book[0] = 116

	// Slice (start, end), 0 based 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was 1/2 price!")

	// Multi line
	poem := `
		The road goes ever on
		Down from the door where it began
		...
		`
	fmt.Println(poem)

}

package main

import (
	"errors"
	"fmt"
)

func main() {
	var input string
	// the Stack is already defined
	var solver Stack

	fmt.Scan(&input)

	for _, char := range input {
		solver.Push(char)
	}

	for {
		char, err := solver.Pop()
		if err != nil {
			break
		}

		fmt.Print(string(char))
	}

	fmt.Println()
}

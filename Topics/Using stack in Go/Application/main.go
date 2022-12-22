package main

import (
	"errors"
	"fmt"
)

func main() {
	var word string
	// the Stack and the Queue are already defined
	var stackSolver Stack
	var queueSolver Queue

	fmt.Scan(&word)

	for _, char := range word {
		stackSolver.Push(char)
		queueSolver.Push(char)
	}

	for {
		a, err1 := stackSolver.Pop()
		b, err2 := queueSolver.Pop()

		if err1 != nil || err2 != nil {
			fmt.Println("Palindrome")
			break
		}

		if a != b {
			fmt.Println("Not palindrome")
			break
		}
	}
}

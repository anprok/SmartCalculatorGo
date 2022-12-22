package main

import (
	"errors"
	"fmt"
)

func main() {
	var brackets string
	// the Queue is already included
	var solver Queue

	fmt.Scan(&brackets)

	for _, bracket := range brackets {
		if bracket == '(' {
			solver.Push(bracket)
		}
		if bracket == ')' {
			_, err := solver.Pop()
			if err != nil {
				fmt.Println("error")
				return
			}
		}
	}

	_, err := solver.Pop()
	if err != nil {
		fmt.Println("All brackets are closed!")
	} else {
		fmt.Println("error")
	}
}

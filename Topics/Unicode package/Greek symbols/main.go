package main

import (
	"fmt"
	"unicode"
)

func main() {
	var word string
	fmt.Scan(&word)

	isGreek := false
	for _, char := range word {
		if unicode.In(char, unicode.Greek) {
			isGreek = true
			break
		}
	}

	if isGreek {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var source string
	fmt.Scan(&source)

	reHeads := regexp.MustCompile(`heads`)
	reTails := regexp.MustCompile(`tails`)

	heads := reHeads.MatchString(source)
	tails := reTails.MatchString(source)

	switch {
	case heads && tails:
		fmt.Println("both")
	case heads:
		fmt.Println("heads")
	case tails:
		fmt.Println("tails")
	default:
		fmt.Println("none")
	}
}

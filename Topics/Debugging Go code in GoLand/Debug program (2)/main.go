package main

import "fmt"

const (
	class  = 36
	pepper = "Pepper"
	no     = "No"
)

func main() {
	var number int
	fmt.Scanln(&number)

	switch number % class {
	case 0:
		fmt.Println(pepper)
	default:
		fmt.Println(no)
	}
}

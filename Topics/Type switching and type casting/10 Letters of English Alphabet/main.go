package main

import "fmt"

func main() {
	letters := [10]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	//put your code here
	for _, letter := range letters {
		fmt.Println(string(letter))
	}
}

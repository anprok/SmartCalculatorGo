package main

import "fmt"

func main() {
	var age interface{} = 13
	switch age.(type) {
	case nil:
		fmt.Println("Age is nil type")
	case int:
		fmt.Println("Age is int type")
	case string:
		fmt.Println("Age is string type")
	default:
		fmt.Println("Type is not defined")
	}
}

package main

import (
	"fmt"
	"log"
	"reflect"
)

func solve(sl1, sl2 []string) []string {
	sl1 = append(sl1, sl2...) // append sl2 to sl1 here!

	return sl1 // return the appended slice here!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	sl1, sl2 := make([]string, num1), make([]string, num2)

	for i := range sl1 {
		fmt.Scanln(&sl1[i])
	}

	for i := range sl2 {
		fmt.Scanln(&sl2[i])
	}

	if !reflect.DeepEqual(append(sl1, sl2...), solve(sl1, sl2)) {
		log.Fatal("slices were not properly appended")
	}
	fmt.Println(solve(sl1, sl2))
}

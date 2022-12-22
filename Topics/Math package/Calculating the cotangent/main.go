package main

import (
	"fmt"
	"math"
)

func main() {
	var radians float64
	fmt.Scanln(&radians)

	// Please use any of the above identities to calculate the 'cotangent' below:
	cotangent := 1 / math.Tan(radians)

	fmt.Printf("%.2f", cotangent)
}

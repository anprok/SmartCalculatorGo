package main

import (
	"fmt"
	"time"
)

func main() {
	var year, month, day int

	// Chicken born time:
	fmt.Scan(&year, &month, &day)
	chicken := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// Egg born time:
	fmt.Scan(&year, &month, &day)
	egg := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	switch {
	case chicken.Before(egg):
		fmt.Println("Chicken is first!")
	case chicken.After(egg):
		fmt.Println("Egg is first!")
	default:
		fmt.Println("Time collapsed")
	}
}

package main

import "fmt"

const (
	maxTax    = 15
	middleTax = 10
	lowTax    = 5

	maxLimit    = 100_000
	middleLimit = 10_000
)

func main() {
	var amount int
	fmt.Scanln(&amount)

	tax := 0
	switch {
	case amount > maxLimit:
		tax += (amount - maxLimit) * maxTax / 100
		amount = maxLimit
		fallthrough
	case amount > middleLimit:
		tax += (amount - middleLimit) * middleTax / 100
		amount = middleLimit
		fallthrough
	default:
		tax += amount * lowTax / 100
	}
	fmt.Println(tax)
}

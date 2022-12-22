package main

import (
	"fmt"
	"math"
)

const ToRadians = math.Pi / 180

func main() {
	var x, y, angle float64

	fmt.Scan(&x, &y, &angle)

	angle *= ToRadians

	// Calculate the new coordinates below:
	newX := x*math.Cos(angle) - y*math.Sin(angle)
	newY := x*math.Sin(angle) + y*math.Cos(angle)

	// Round newX and newY to the nearest integer:
	newX = math.Round(newX)
	newY = math.Round(newY)

	fmt.Println(newX, newY)
}

package main

import (
	"fmt"
	"time"
)

const (
	Year  = 1902
	Month = 11
	Day   = 2
)

// DO NOT delete the comments on each line below!
func main() {
	utc := time.Date(Year, Month, Day, 0, 0, 0, 0, time.UTC)
	london := utc.Add(1 * time.Hour)                // nolint: gomnd
	moscow := utc.Add(3 * time.Hour)                // nolint: gomnd
	newYork := utc.Add(-4 * time.Hour)              // nolint: gomnd
	hongKong := utc.Add(8 * time.Hour)              // nolint: gomnd
	tokyo := utc.Add(9 * time.Hour)                 // nolint: gomnd
	darwin := utc.Add(9*time.Hour + 30*time.Minute) // nolint: gomnd

	fmt.Printf("London: %v\n", london)
	fmt.Printf("Moscow: %v\n", moscow)
	fmt.Printf("New York: %v\n", newYork)
	fmt.Printf("Hong Kong: %v\n", hongKong)
	fmt.Printf("Tokyo: %v\n", tokyo)
	fmt.Printf("Darwin: %v\n", darwin)
}

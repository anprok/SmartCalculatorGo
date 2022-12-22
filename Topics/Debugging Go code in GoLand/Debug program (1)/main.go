package main

import "fmt"

func main() {
	freezing := 1999
	defrosting := 2999

	leapYear := 0
	for year := freezing; year <= defrosting; year++ {
		if isLeap(year) {
			leapYear++
		}
	}
	// please, don't correct this output
	fmt.Printf("%d leap years from %d to %d", leapYear, freezing, defrosting)
}

func isLeap(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	nextYear := 2022 + 1

	firstJanuary := time.Date(nextYear, time.January, 1, 0, 0, 0, 0, time.Local)

	// The hidden 'check' function checks your answer:
	check(firstJanuary.Sub(now))
}

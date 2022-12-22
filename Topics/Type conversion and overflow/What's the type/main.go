package main

import "fmt"

const (
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func main() {
	var source float64
	fmt.Scan(&source)

	// check if the source value can be converted to int8 value without overflow
	if source < float64(MaxUint8) {
		fmt.Println("uint8")
	}
	if source < float64(MaxUint16) {
		fmt.Println("uint16")
	}
	if source < float64(MaxUint32) {
		fmt.Println("uint32")
	}
	if source < float64(MaxUint64) {
		fmt.Println("uint64")
	}
	// add the same check for 16, 32, 64 bits...

	// check if we the source value is out of the range for all types
	if float64(MaxUint64) < source {
		fmt.Println("unsupported type")
	}
}

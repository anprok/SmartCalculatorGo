package main

import (
	"fmt"
	"strings"
)

const quote = "Knowledge is better than money"

func main() {
	reader := strings.NewReader(quote)

	buf := make([]byte, 5)
	reader.ReadAt(buf, 25)
	fmt.Println(string(buf))
}

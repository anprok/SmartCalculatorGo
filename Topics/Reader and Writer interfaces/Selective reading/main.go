package main

import (
	"fmt"
	"io"
	"strings"
)

const textSource = "Supercalifragilisticexpialidocious"

func main() {
	fmt.Println("full text:", textSource)

	reader := strings.NewReader(textSource)

	// read 7 bytes from reader
	p := make([]byte, 7)
	reader.Seek(-7, io.SeekEnd)
	if _, err := reader.Read(p); err != nil {
		fmt.Println(err)
	}
	fmt.Print("selected: ", string(p), " ")
}

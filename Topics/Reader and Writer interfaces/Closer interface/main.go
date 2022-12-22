package main

import (
	"fmt"
	"io"
)

type FunnyBox struct {
}

func (s *FunnyBox) Close() error {
	fmt.Println("Bey!")
	return nil
}

func main() {
	var c io.Closer
	c = &FunnyBox{}
	c.Close()
}

package main

import (
	"fmt"
	"io"
)

type WisdomBox struct {
	quotes string
}

func (wb *WisdomBox) Write(p []byte) (n int, err error) {
	wb.quotes += string(p)
	return len(p), nil
}

func (wb *WisdomBox) Read(p []byte) (n int, err error) {
	p = []byte(wb.quotes)
	return len(p), nil
}

func main() {
	var box io.ReadWriter = &WisdomBox{}
	fmt.Fprintln(box, "Never underestimate anyone")
	fmt.Fprintln(box, "Home is where the heart is")
	fmt.Fprintln(box, "You can be whoever you want to be")

	p := make([]byte, 1024)
	n, _ := box.Read(p)
	fmt.Print(string(p[:n-1]))
}

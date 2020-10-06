package main

import "golang.org/x/tour/reader"

// MyReader defintion
type MyReader struct{}

// Implementing the io.Reader interface for MyReader
func (r MyReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = byte('A')
	}
	return len(buf), nil
}

func runEx7() {
	reader.Validate(MyReader{})
}

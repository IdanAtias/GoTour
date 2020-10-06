package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

var numLetters byte = byte('Z') - byte('A') + 1

func (r13 *rot13Reader) Read(buf []byte) (int, error) {
	n, err := r13.r.Read(buf)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		if !unicode.IsLetter(rune(buf[i])) {
			continue
		}
		startPoint := byte('a')
		if unicode.IsUpper(rune(buf[i])) {
			startPoint = byte('A')
		}
		distFromStart := buf[i] - startPoint
		newDistFromStart := (distFromStart + 13) % numLetters
		buf[i] = startPoint + newDistFromStart
	}
	return n, err
}

func runEx8() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

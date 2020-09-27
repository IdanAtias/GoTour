package main

import "golang.org/x/tour/pic"

// Pic - build pic
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}
	for n := 0; n < dy; n++ {
		for m := 0; m < dx; m++ {
			pic[n][m] = uint8(n * m)
		}
	}
	return pic
}

func runEx2() {
	pic.Show(Pic)
}

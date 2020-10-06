package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image type defintion
type Image struct{}

// ColorModel use RGBAModel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds define image boundaries
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

// At set the color for each pixle
func (i Image) At(x, y int) color.Color {
	v := uint8((x * y) % 256) // could be also something else
	return color.RGBA{v, v, 255, 255}
}

func runEx9() {
	m := Image{}
	pic.ShowImage(m)
}

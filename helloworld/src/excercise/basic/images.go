package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h, v int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.ColorModel {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA(uint32(i.v), uint32(i.v), 255, 255)
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

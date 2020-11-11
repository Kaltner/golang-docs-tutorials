package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
	v uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.v + uint8(x), i.v + uint8(y), 255, 255}
}

func main() {
	m := Image{
		w: 100,
		h: 10,
		v: 16,
	}
	pic.ShowImage(m)
}

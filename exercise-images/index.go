package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	MIN, MAX image.Point
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.MIN.X, i.MIN.Y, i.MAX.X, i.MAX.Y)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{image.Point{0, 0}, image.Point{100, 100}}
	pic.ShowImage(m)
}

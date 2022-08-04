package main

import (
	"golang.org/x/tour/pic" 
	"image"
	"image/color"
) 

type Image struct{ w, h int }

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (_ Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (_ Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{1000, 1000}
	pic.ShowImage(m)
}
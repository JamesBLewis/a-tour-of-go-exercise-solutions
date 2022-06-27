package main

import (
"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for b := range a {
		a[b] = make([]uint8, dx)
	}
	return a
}

func main() {
	pic.Show(Pic)
}

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dx)
	y := make([]uint8, dy)
	for i := range x {
		for j := range y {
			y[j] = uint8(i * j / 2 * i)
		}
		x[i] = y
	}
	return x
}

func main() {
	pic.Show(Pic)
}

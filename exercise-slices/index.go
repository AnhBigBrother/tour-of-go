package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := [][]uint8{}
	for i := 0; i < dy; i++ {
		p = append(p, make([]uint8, dx))
	}
	for i := range p {
		for j := range p[i] {
			p[i][j] = uint8(i ^ j)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}

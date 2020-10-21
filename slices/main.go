package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)
	for x := range picture {
		picture[x] = make([]uint8, dy)
		for y := range picture {
			picture[x][y] = uint8((x + 1) * (y + 1))
		}
	}

	return picture
}

func main() {
	fmt.Println(Pic(2, 2))
}

package main

import "fmt"

const serial, size = 7803, 300

func main() {
	var cells [size + 1][size + 1]int
	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			cells[y][x] = ((x+10)*y+serial)*(x+10)/100%10 - 5 + cells[y-1][x] + cells[y][x-1] - cells[y-1][x-1]
		}
	}

	var p1m, p1x, p1y int
	var p2m, p2x, p2y, p2s int
	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			for s := 1; x+s <= size && y+s <= size; s++ {
				sum := cells[y+s-1][x+s-1] + cells[y-1][x-1] - cells[y-1][x+s-1] - cells[y+s-1][x-1]
				if s == 3 && sum > p1m {
					p1m, p1x, p1y = sum, x, y
				}
				if sum > p2m {
					p2m, p2x, p2y, p2s = sum, x, y, s
				}
			}
		}
	}
	fmt.Printf("%d,%d\n", p1x, p1y)
	fmt.Printf("%d,%d,%d\n", p2x, p2y, p2s)
}

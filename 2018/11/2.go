package main

import "fmt"

const serial, size = 7803, 300

func main() {
	var cells [size + 1][size + 1]int
	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			cells[x][y] = ((x+10)*y+serial)*(x+10)/100%10 - 5 + cells[x][y-1] + cells[x-1][y] - cells[x-1][y-1]
		}
	}

	var max, mx, my, ms int
	for x := 1; x <= size-2; x++ {
		for y := 1; y <= size-2; y++ {
			sum := cells[x+2][y+2] + cells[x-1][y-1] - cells[x+2][y-1] - cells[x-1][y+2]
			if sum > max {
				max, mx, my = sum, x, y
			}
		}
	}
	fmt.Printf("%d,%d\n", mx, my)

	max = 0
	for s := 1; s <= size; s++ {
		for x := 1; x <= size-s+1; x++ {
			for y := 1; y <= size-s+1; y++ {
				sum := cells[x+s-1][y+s-1] + cells[x-1][y-1] - cells[x+s-1][y-1] - cells[x-1][y+s-1]
				if sum > max {
					max, mx, my, ms = sum, x, y, s
				}
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", mx, my, ms)
}

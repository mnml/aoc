package main

import (
	"fmt"
	"image"
	"maps"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
	}

	grid := map[image.Point]int{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '@' {
				grid[image.Point{x, y}] = 1
			}
		}
	}

	removed := 0
	for i := 0; ; i++ {
		next := maps.Clone(grid)
		for p := range grid {
			rolls := 0
			for _, d := range delta {
				rolls += grid[p.Add(d)]
			}
			if rolls < 4 {
				delete(next, p)
				removed++
			}
		}
		if i == 0 {
			fmt.Println(removed)
		}
		if maps.Equal(next, grid) {
			fmt.Println(removed)
			break
		}
		grid = next
	}
}

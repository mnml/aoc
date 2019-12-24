package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	grid := map[image.Point]int{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = map[rune]int{'.': 0, '#': 1}[r]
		}
	}

	seen := map[int]struct{}{}
	for {
		bio := 0

		next := map[image.Point]int{}
		for p, r := range grid {
			adj := 0
			for _, d := range []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
				adj += grid[p.Add(d)]
			}

			next[p] = 0
			if adj == 1 || r == 0 && adj == 2 {
				next[p] = 1
				bio += 1 << (5*p.Y + p.X)
			}
		}
		grid = next

		if _, ok := seen[bio]; ok {
			fmt.Println(bio)
			return
		}
		seen[bio] = struct{}{}
	}
}

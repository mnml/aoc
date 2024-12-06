package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			if r == '^' {
				start = image.Point{x, y}
			}
			grid[image.Point{x, y}] = r
		}
	}

	patrol := func(o image.Point) int {
		delta := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		p, d, seen := start, 0, map[image.Point]int{}
		for {
			if _, ok := grid[p]; !ok {
				return len(seen)
			} else if 1<<d&seen[p] != 0 {
				return -1
			}
			seen[p] |= 1 << d
			if n := p.Add(delta[d]); grid[n] == '#' || n == o {
				d = (d + 1) % len(delta)
			} else {
				p = n
			}
		}
	}

	part2 := 0
	for p := range grid {
		if patrol(p) == -1 {
			part2++
		}
	}
	fmt.Println(patrol(image.Point{-1, -1}))
	fmt.Println(part2)
}

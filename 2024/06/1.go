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
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '^' {
				start = image.Point{x, y}
			}
			grid[image.Point{x, y}] = r
		}
	}

	patrol := func(o image.Point) map[image.Point]int {
		delta := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		p, d, seen := start, 0, map[image.Point]int{}
		for {
			if _, ok := grid[p]; !ok {
				return seen
			} else if 1<<d&seen[p] != 0 {
				return nil
			}
			seen[p] |= 1 << d
			if n := p.Add(delta[d]); grid[n] == '#' || n == o {
				d = (d + 1) % len(delta)
			} else {
				p = n
			}
		}
	}

	part1, part2 := patrol(image.Point{-1, -1}), 0
	for p := range part1 {
		if patrol(p) == nil {
			part2++
		}
	}
	fmt.Println(len(part1))
	fmt.Println(part2)
}

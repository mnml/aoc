package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	grid := map[image.Point]rune{}
	var start image.Point
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}

	var beam func(image.Point) int
	cache := map[image.Point]int{}
	part1 := map[image.Point]struct{}{}

	beam = func(p image.Point) (n int) {
		if n, ok := cache[p]; ok {
			return n
		}
		defer func() { cache[p] = n }()

		p.Y++
		if _, ok := grid[p]; !ok {
			return 1
		}
		if grid[p] == '^' {
			part1[p] = struct{}{}
			return beam(image.Point{p.X - 1, p.Y}) + beam(image.Point{p.X + 1, p.Y})
		}
		return beam(p)
	}

	part2 := beam(start)
	fmt.Println(len(part1))
	fmt.Println(part2)
}

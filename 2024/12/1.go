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
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}

	seen := map[image.Point]bool{}
	part1, part2 := 0, 0
	for p := range grid {
		if seen[p] {
			continue
		}
		seen[p] = true

		area := 1
		perimeter, sides := 0, 0
		queue := []image.Point{p}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				if n := p.Add(d); grid[n] != grid[p] {
					perimeter++
					r := p.Add(image.Point{-d.Y, d.X})
					if grid[r] != grid[p] || grid[r.Add(d)] == grid[p] {
						sides++
					}
				} else if !seen[n] {
					seen[n] = true
					queue = append(queue, n)
					area++
				}
			}
		}
		part1 += area * perimeter
		part2 += area * sides
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

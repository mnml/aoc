package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	grid, p := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == 'S' {
				p = image.Point{x, y}
			}
			grid[image.Point{x, y}] = r
		}
	}

	dist := map[image.Point]int{p: 0}
	for grid[p] != 'E' {
		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := p.Add(d)
			if _, ok := dist[n]; !ok && grid[n] != '#' {
				p, dist[n] = n, dist[p]+1
			}
		}
	}

	part1, part2 := 0, 0
	for p := range dist {
		for q := range dist {
			d := abs(q.X-p.X) + abs(q.Y-p.Y)
			if d <= 20 && dist[q]-dist[p]-d >= 100 {
				if d <= 2 {
					part1++
				}
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

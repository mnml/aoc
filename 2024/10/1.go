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

	part1, part2 := 0, 0
	for p := range grid {
		if grid[p] == '0' {
			part1 += dfs(grid, p, map[image.Point]bool{})
			part2 += dfs(grid, p, nil)
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func dfs(grid map[image.Point]rune, p image.Point, seen map[image.Point]bool) (score int) {
	if grid[p] == '9' {
		if seen[p] {
			return 0
		} else if seen != nil {
			seen[p] = true
		}
		return 1
	}
	for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		if n := p.Add(d); grid[n] == grid[p]+1 {
			score += dfs(grid, n, seen)
		}
	}
	return score
}

package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type State struct {
	P image.Point
	D int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	grid := map[image.Point]rune{}
	start, end := image.Point{1, 0}, image.Point{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
			end = image.Point{x - 1, y}
		}
	}

	delta := map[rune]image.Point{'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0}}
	seen := map[image.Point]struct{}{}

	var dfs func(image.Point, int, bool) int
	dfs = func(start image.Point, dist int, slopes bool) (best int) {
		if start == end {
			return dist
		}
		seen[start] = struct{}{}
		for _, d := range delta {
			p := start.Add(d)
			if n, ok := delta[grid[start]]; slopes && ok && n != d {
				continue
			}
			if _, ok := seen[p]; !ok && grid[p] != '#' && grid[p] != '\x00' {
				best = max(best, dfs(p, dist+1, slopes))
			}
		}
		delete(seen, start)
		return best
	}

	fmt.Println(dfs(start, 0, true))
	clear(seen)
	fmt.Println(dfs(start, 0, false))
}

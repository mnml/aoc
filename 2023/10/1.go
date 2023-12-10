package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}

	grid[start] = map[[4]bool]rune{
		{true, false, true, false}: '|', {false, true, false, true}: '-',
		{true, true, false, false}: 'L', {true, false, false, true}: 'J',
		{false, false, true, true}: '7', {false, true, true, false}: 'F',
	}[[4]bool{
		strings.ContainsRune("7F|", grid[start.Add(image.Point{0, -1})]),
		strings.ContainsRune("-7J", grid[start.Add(image.Point{1, 0})]),
		strings.ContainsRune("JL|", grid[start.Add(image.Point{0, 1})]),
		strings.ContainsRune("-FL", grid[start.Add(image.Point{-1, 0})]),
	}]

	path, area := []image.Point{}, 0
	for p, n := start, start; p == start || n != start; path = append(path, p) {
		p, n = n, start

		for _, d := range map[rune][]image.Point{
			'|': {{0, -1}, {0, 1}}, '-': {{1, 0}, {-1, 0}}, 'L': {{0, -1}, {1, 0}},
			'J': {{0, -1}, {-1, 0}}, '7': {{0, 1}, {-1, 0}}, 'F': {{0, 1}, {1, 0}},
		}[grid[p]] {
			if !slices.Contains(path, p.Add(d)) {
				n = p.Add(d)
			}
		}

		area += p.X*n.Y - p.Y*n.X
	}

	fmt.Println(len(path) / 2)
	fmt.Println((int(math.Abs(float64(area)))-len(path))/2 + 1)
}

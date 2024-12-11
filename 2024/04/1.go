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

	adj := func(p image.Point, l int) []string {
		delta := []image.Point{
			{0, -1}, {1, 0}, {0, 1}, {-1, 0},
			{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
		}

		words := make([]string, len(delta))
		for i, d := range delta {
			for n := range l {
				words[i] += string(grid[p.Add(d.Mul(n))])
			}
		}
		return words
	}

	part1, part2 := 0, 0
	for p := range grid {
		part1 += strings.Count(strings.Join(adj(p, 4), " "), "XMAS")
		part2 += strings.Count("AMAMASASAMAMAS", strings.Join(adj(p, 2)[4:], ""))
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

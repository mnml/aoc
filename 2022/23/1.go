package main

import (
	"fmt"
	"image"
	"os"
	"reflect"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	grid := map[image.Point]struct{}{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '#' {
				grid[image.Point{x, y}] = struct{}{}
			}
		}
	}

	sides := [][]image.Point{
		{{0, -1}, {1, -1}, {-1, -1}},
		{{0, 1}, {1, 1}, {-1, 1}},
		{{-1, 0}, {-1, -1}, {-1, 1}},
		{{1, 0}, {1, -1}, {1, 1}},
	}

	for i := 0; ; i++ {
		prop := map[image.Point]image.Point{}
		count := map[image.Point]int{}

		for p := range grid {
			neigh := map[int]int{}
			for i := range sides {
				for _, q := range sides[i] {
					if _, ok := grid[p.Add(q)]; ok {
						neigh[i]++
					}
				}
			}

			if len(neigh) == 0 {
				continue
			}

			for d := 0; d < len(sides); d++ {
				if dir := (i + d) % len(sides); neigh[dir] == 0 {
					prop[p] = p.Add(sides[dir][0])
					count[prop[p]]++
					break
				}
			}
		}

		newGrid := map[image.Point]struct{}{}
		for p := range grid {
			if _, ok := prop[p]; ok && count[prop[p]] == 1 {
				p = prop[p]
			}
			newGrid[p] = struct{}{}
		}

		if i == 9 {
			var r image.Rectangle
			for p := range newGrid {
				r = r.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
			}
			fmt.Println(r.Dx()*r.Dy() - len(newGrid))
		}
		if reflect.DeepEqual(grid, newGrid) {
			fmt.Println(i + 1)
			break
		}

		grid = newGrid
	}
}

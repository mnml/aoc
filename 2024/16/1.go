package main

import (
	"cmp"
	"fmt"
	"image"
	"maps"
	"math"
	"os"
	"slices"
	"strings"
)

type State struct {
	Pos image.Point
	Dir image.Point
}

type QI struct {
	State State
	Cost  int
	Path  map[image.Point]struct{}
}

func main() {
	input, _ := os.ReadFile("input.txt")

	var start image.Point
	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == 'S' {
				start = image.Point{x, y}
			}
			grid[image.Point{x, y}] = r
		}
	}

	dist := map[State]int{}
	queue := []QI{{
		State{start, image.Point{1, 0}},
		0,
		map[image.Point]struct{}{start: {}},
	}}

	part1, part2 := math.MaxInt, map[image.Point]struct{}{}
	for len(queue) > 0 {
		slices.SortFunc(queue, func(a, b QI) int {
			return cmp.Compare(a.Cost, b.Cost)
		})
		i := queue[0]
		queue = queue[1:]

		if c, ok := dist[i.State]; ok && c < i.Cost {
			continue
		}
		dist[i.State] = i.Cost

		if grid[i.State.Pos] == 'E' && i.Cost <= part1 {
			part1 = i.Cost
			maps.Copy(part2, i.Path)
		}

		for d, c := range map[image.Point]int{
			i.State.Dir:                     1,
			{-i.State.Dir.Y, i.State.Dir.X}: 1001,
			{i.State.Dir.Y, -i.State.Dir.X}: 1001,
		} {
			n := State{i.State.Pos.Add(d), d}
			if grid[n.Pos] == '#' {
				continue
			}
			path := maps.Clone(i.Path)
			path[n.Pos] = struct{}{}
			queue = append(queue, QI{n, i.Cost + c, path})
		}
	}
	fmt.Println(part1)
	fmt.Println(len(part2))
}

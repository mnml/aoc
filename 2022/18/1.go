package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y, p.Z + q.Z}
}

func main() {
	input, _ := os.ReadFile("input.txt")

	lava := map[Point]struct{}{}
	min := Point{math.MaxInt, math.MaxInt, math.MaxInt}
	max := Point{math.MinInt, math.MinInt, math.MinInt}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var p Point
		fmt.Sscanf(s, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		lava[p] = struct{}{}

		min = Point{Min(min.X, p.X), Min(min.Y, p.Y), Min(min.Z, p.Z)}
		max = Point{Max(max.X, p.X), Max(max.Y, p.Y), Max(max.Z, p.Z)}
	}
	min = min.Add(Point{-1, -1, -1})
	max = max.Add(Point{1, 1, 1})

	delta := []Point{
		{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	}

	part1 := 0
	for p := range lava {
		for _, d := range delta {
			if _, ok := lava[p.Add(d)]; !ok {
				part1++
			}
		}
	}
	fmt.Println(part1)

	queue := []Point{min}
	visited := map[Point]struct{}{min: {}}

	part2 := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range delta {
			next := cur.Add(d)

			if _, ok := lava[next]; ok {
				part2++
			} else if _, ok := visited[next]; !ok &&
				next.X >= min.X && next.X <= max.X &&
				next.Y >= min.Y && next.Y <= max.Y &&
				next.Z >= min.Z && next.Z <= max.Z {
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}
	fmt.Println(part2)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

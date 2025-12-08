package main

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func dist2(p, q Point) int {
	dx, dy, dz := q.X-p.X, q.Y-p.Y, q.Z-p.Z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	input, _ := os.ReadFile("input.txt")

	pairs, circuits := [][2]Point{}, []map[Point]bool{}
	for _, s := range strings.Fields(string(input)) {
		var b Point
		fmt.Sscanf(s, "%d,%d,%d", &b.X, &b.Y, &b.Z)
		for _, c := range circuits {
			pairs = append(pairs, [2]Point{b, slices.Collect(maps.Keys(c))[0]})
		}
		circuits = append(circuits, map[Point]bool{b: true})
	}
	slices.SortFunc(pairs, func(a, b [2]Point) int {
		return cmp.Compare(dist2(a[0], a[1]), dist2(b[0], b[1]))
	})

	for i, p := range pairs {
		c1 := slices.IndexFunc(circuits, func(c map[Point]bool) bool { return c[p[0]] })
		c2 := slices.IndexFunc(circuits, func(c map[Point]bool) bool { return c[p[1]] })
		if c1 != c2 {
			maps.Copy(circuits[c1], circuits[c2])
			circuits = slices.Delete(circuits, c2, c2+1)
		}

		if i+1 == 1000 {
			slices.SortFunc(circuits, func(a, b map[Point]bool) int {
				return -cmp.Compare(len(a), len(b))
			})
			fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
		}
		if len(circuits) == 1 {
			fmt.Println(p[0].X * p[1].X)
			break
		}
	}
}

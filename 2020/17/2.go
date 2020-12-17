package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point [4]int

func (p Point) Add(q Point) Point {
	for i := range p {
		p[i] += q[i]
	}
	return p
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	grid := map[Point]struct{}{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '#' {
				grid[Point{x, y}] = struct{}{}
			}
		}
	}

	fmt.Println(run(grid, 3, 6))
	fmt.Println(run(grid, 4, 6))
}

func run(grid map[Point]struct{}, dim, cycles int) int {
	for i := 0; i < cycles; i++ {
		neigh := map[Point]int{}
		for p := range grid {
			for _, d := range delta(dim)[1:] {
				neigh[p.Add(d)]++
			}
		}

		new := map[Point]struct{}{}
		for p, n := range neigh {
			if _, ok := grid[p]; ok && n == 2 || n == 3 {
				new[p] = struct{}{}
			}
		}
		grid = new
	}
	return len(grid)
}

func delta(dim int) (ds []Point) {
	if dim == 0 {
		return []Point{{}}
	}
	for _, v := range []int{0, 1, -1} {
		for _, p := range delta(dim - 1) {
			p[dim-1] = v
			ds = append(ds, p)
		}
	}
	return
}

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

	grid := map[Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[Point{x, y}] = r
		}
	}

	fmt.Println(run(grid, 3, 6))
	fmt.Println(run(grid, 4, 6))
}

func run(grid map[Point]rune, dim, cycles int) (sum int) {
	delta := delta(dim)[1:]

	for i := 0; i < cycles; i++ {
		for p := range grid {
			for _, d := range delta {
				grid[p.Add(d)] = grid[p.Add(d)]
			}
		}

		new := map[Point]rune{}
		for p, r := range grid {
			neigh := 0
			for _, d := range delta {
				if grid[p.Add(d)] == '#' {
					neigh++
				}
			}

			if r == '#' && neigh == 2 || neigh == 3 {
				new[p] = '#'
			}
		}
		grid = new
	}

	for _, r := range grid {
		if r == '#' {
			sum++
		}
	}
	return
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

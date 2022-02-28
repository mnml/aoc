package main

import (
	"fmt"
	"image"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	delta := []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	grid := map[image.Point]int{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = int(r - '0')
		}
	}

	part1, part2 := 0, []int{}
out:
	for p, h := range grid {
		for _, d := range delta {
			if a, ok := grid[p.Add(d)]; ok && a <= h {
				continue out
			}
		}
		part1 += h + 1

		queue, size := []image.Point{p}, 0
		grid[p] = 9
		for len(queue) > 0 {
			p, queue = queue[0], queue[1:]
			size++

			for _, d := range delta {
				p := p.Add(d)
				if h, ok := grid[p]; !ok || h == 9 {
					continue
				}

				queue = append(queue, p)
				grid[p] = 9
			}
		}
		part2 = append(part2, size)
	}

	fmt.Println(part1)
	sort.Sort(sort.Reverse(sort.IntSlice(part2)))
	fmt.Println(part2[0] * part2[1] * part2[2])
}

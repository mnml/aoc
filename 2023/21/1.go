package main

import (
	"fmt"
	"image"
	"maps"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	grid, start, bounds := map[image.Point]rune{}, image.Point{}, image.Rectangle{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			p := image.Point{x, y}
			if r == 'S' {
				start = p
				r = '.'
			}
			grid[p] = r
			bounds = bounds.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
		}
	}
	w := bounds.Dx()

	garden := map[image.Point]struct{}{start: {}}
	seen := map[image.Point]struct{}{start: {}}
	plots := map[int]int{0: 1}

	var s int
loop:
	for s = 1; ; s++ {
		ng := map[image.Point]struct{}{}
		for p := range garden {
			for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				q := p.Add(d)
				if _, ok := seen[q]; !ok && grid[q.Mod(bounds)] == '.' {
					ng[q] = struct{}{}
				}
			}
		}
		maps.Copy(seen, ng)
		plots[s] = plots[s-2] + len(ng)
		garden = ng

		for i := s; i > s-w; i-- {
			d1 := (plots[i] - plots[i-2]) - (plots[i-w] - plots[i-w-2])
			d2 := (plots[i-w] - plots[i-w-2]) - (plots[i-2*w] - plots[i-2*w-2])
			if d1 != d2 {
				continue loop
			}
		}
		break
	}

	for ; s <= 26501365; s++ {
		plots[s] = plots[s-2] + 2*plots[s-w] - 2*plots[s-w-2] - plots[s-2*w] + plots[s-2*w-2]
	}

	fmt.Println(plots[64])
	fmt.Println(plots[26501365])
}

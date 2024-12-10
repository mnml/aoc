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
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}

	bfs := func(all bool) (score int) {
		for start := range grid {
			if grid[start] != '0' {
				continue
			}

			queue, seen := []image.Point{start}, map[image.Point]struct{}{start: {}}
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]

				if grid[p] == '9' {
					score++
					continue
				}

				for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
					n := p.Add(d)
					if _, ok := seen[n]; grid[n] == grid[p]+1 && (!ok || all) {
						queue, seen[n] = append(queue, n), struct{}{}
					}
				}
			}
		}
		return score
	}

	fmt.Println(bfs(false))
	fmt.Println(bfs(true))
}

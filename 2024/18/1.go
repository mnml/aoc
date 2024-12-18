package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	size := 70

	bytes := []image.Point{}
	for _, s := range strings.Fields(string(input)) {
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		bytes = append(bytes, image.Point{x, y})
	}

	grid := map[image.Point]bool{}
	for y := range size + 1 {
		for x := range size + 1 {
			grid[image.Point{x, y}] = true
		}
	}

loop:
	for b := range bytes {
		grid[bytes[b]] = false

		queue, dist := []image.Point{{0, 0}}, map[image.Point]int{{0, 0}: 0}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p == (image.Point{size, size}) {
				if b == 1024 {
					fmt.Println(dist[p])
				}
				continue loop
			}

			for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				n := p.Add(d)
				if _, ok := dist[n]; !ok && grid[n] {
					queue, dist[n] = append(queue, n), dist[p]+1
				}
			}
		}
		fmt.Printf("%d,%d\n", bytes[b].X, bytes[b].Y)
		break
	}
}

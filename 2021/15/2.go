package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fields := strings.Fields(strings.TrimSpace(string(input)))
	w, h := len(fields[0]), len(fields)

	cave1, cave2 := map[image.Point]int{}, map[image.Point]int{}
	for y, s := range fields {
		for x, r := range s {
			v := int(r - '0')
			cave1[image.Point{x, y}] = v
			for j := 0; j < 5; j++ {
				for i := 0; i < 5; i++ {
					cave2[image.Point{i*w + x, j*h + y}] = (v+i+j-1)%9 + 1
				}
			}
		}
	}
	fmt.Println(dist(image.Point{w - 1, h - 1}, cave1))
	fmt.Println(dist(image.Point{w*5 - 1, h*5 - 1}, cave2))
}

func dist(end image.Point, cave map[image.Point]int) int {
	queue := map[image.Point]struct{}{{0, 0}: {}}
	dist := map[image.Point]int{{0, 0}: 0}

	for len(queue) > 0 {
		var cur *image.Point
		for p := range queue {
			if cur == nil || dist[p] < dist[*cur] {
				cur = &image.Point{p.X, p.Y}
			}
		}

		if *cur == end {
			return dist[*cur]
		}
		delete(queue, *cur)

		for _, d := range []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			p := cur.Add(d)
			if _, ok := cave[p]; !ok {
				continue
			}

			alt := dist[*cur] + cave[p]
			if _, ok := dist[p]; !ok || alt < dist[p] {
				dist[p] = alt
				queue[p] = struct{}{}
			}
		}
	}
	return -1
}

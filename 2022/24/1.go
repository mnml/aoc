package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type State struct {
	P image.Point
	T int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	vall := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			vall[image.Point{x, y}] = r
		}
	}

	var bliz image.Rectangle
	for p := range vall {
		bliz = bliz.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
	}
	bliz.Min, bliz.Max = bliz.Min.Add(image.Point{1, 1}), bliz.Max.Sub(image.Point{1, 1})

	bfs := func(start image.Point, end image.Point, time int) int {
		delta := map[rune]image.Point{
			'#': {0, 0}, '^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		}

		queue := []State{{start, time}}
		seen := map[State]struct{}{queue[0]: {}}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

		loop:
			for _, d := range delta {
				next := State{cur.P.Add(d), cur.T + 1}
				if next.P == end {
					return next.T
				}

				if _, ok := seen[next]; ok {
					continue
				}
				if r, ok := vall[next.P]; !ok || r == '#' {
					continue
				}

				if next.P.In(bliz) {
					for r, d := range delta {
						if vall[next.P.Sub(d.Mul(next.T)).Mod(bliz)] == r {
							continue loop
						}
					}
				}

				seen[next] = struct{}{}
				queue = append(queue, next)
			}
		}
		return -1
	}

	start, end := bliz.Min.Sub(image.Point{0, 1}), bliz.Max.Sub(image.Point{1, 0})
	fmt.Println(bfs(start, end, 0))
	fmt.Println(bfs(start, end, bfs(end, start, bfs(start, end, 0))))
}

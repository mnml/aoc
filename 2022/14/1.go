package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	rock, maxy := map[image.Point]struct{}{}, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		s := strings.Split(s, " -> ")

		for i := 0; i < len(s)-1; i++ {
			var p, q image.Point
			fmt.Sscanf(s[i], "%d,%d", &p.X, &p.Y)
			fmt.Sscanf(s[i+1], "%d,%d", &q.X, &q.Y)

			for d := (image.Point{sgn(q.X - p.X), sgn(q.Y - p.Y)}); p != q.Add(d); p = p.Add(d) {
				rock[p] = struct{}{}
				if p.Y > maxy {
					maxy = p.Y
				}
			}
		}
	}

	d := []image.Point{{0, 1}, {-1, 1}, {1, 1}}

	part1, part2 := (*int)(nil), 0
	for {
		p := image.Point{500, 0}

		for i := 0; i < len(d); i++ {
			if _, ok := rock[p.Add(d[i])]; !ok && p.Add(d[i]).Y < maxy+2 {
				p = p.Add(d[i])
				if c := part2; part1 == nil && p.Y >= maxy {
					part1 = &c
				}
				i = -1
			}
		}

		rock[p] = struct{}{}
		part2++
		if p.Y == 0 {
			break
		}
	}
	fmt.Println(*part1)
	fmt.Println(part2)
}

func sgn(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}

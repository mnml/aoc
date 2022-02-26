package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	ortho, diag := map[image.Point]int{}, map[image.Point]int{}
	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var p, q image.Point
		fmt.Sscanf(s, "%d,%d -> %d,%d", &p.X, &p.Y, &q.X, &q.Y)

		for d := (image.Point{sgn(q.X - p.X), sgn(q.Y - p.Y)}); p != q.Add(d); p = p.Add(d) {
			if diag[p]++; diag[p] == 2 {
				part2++
			}
			if d.X != 0 && d.Y != 0 {
				continue
			}
			if ortho[p]++; ortho[p] == 2 {
				part1++
			}
		}
	}
	fmt.Println(part1)
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

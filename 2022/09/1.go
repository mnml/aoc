package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	dirs := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}
	rope := make([]image.Point, 10)

	part1, part2 := map[image.Point]struct{}{}, map[image.Point]struct{}{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var dir rune
		var steps int
		fmt.Sscanf(s, "%c %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].Add(dirs[dir])

			for i := 1; i < len(rope); i++ {
				if d := rope[i-1].Sub(rope[i]); abs(d.X) > 1 || abs(d.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{sgn(d.X), sgn(d.Y)})
				}
			}

			part1[rope[1]], part2[rope[len(rope)-1]] = struct{}{}, struct{}{}
		}
	}
	fmt.Println(len(part1))
	fmt.Println(len(part2))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

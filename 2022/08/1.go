package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	trees := map[image.Point]int{}
	for y, s := range strings.Fields(strings.TrimSpace(string(input))) {
		for x, r := range s {
			trees[image.Point{x, y}] = int(r - '0')
		}
	}

	part1, part2 := 0, 0
	for p := range trees {
		vis, score := 0, 1

		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			for np, i := p.Add(d), 0; ; np, i = np.Add(d), i+1 {
				if _, ok := trees[np]; !ok {
					vis, score = 1, score*i
					break
				}

				if trees[np] >= trees[p] {
					score *= i + 1
					break
				}
			}
		}

		part1 += vis
		if score > part2 {
			part2 = score
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

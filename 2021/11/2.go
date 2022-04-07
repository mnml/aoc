package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	octo := map[image.Point]int{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			octo[image.Point{x, y}] = int(r - '0')
		}
	}

	part1, part2 := 0, 0
	for flashes := 0; flashes < len(octo); {
		for p := range octo {
			octo[p]++
		}

		flashes = 0
	loop:
		for p := range octo {
			if octo[p] > 9 {
				flashes += flash(octo, p)
				goto loop
			}
		}

		if part2++; part2 <= 100 {
			part1 += flashes
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func flash(octo map[image.Point]int, p image.Point) int {
	delta := []image.Point{
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
		{0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	octo[p] = 0
	flashes := 1

	for _, d := range delta {
		n := p.Add(d)
		if octo[n] != 0 {
			octo[n]++
		}
		if octo[n] > 9 {
			flashes += flash(octo, n)
		}
	}
	return flashes
}

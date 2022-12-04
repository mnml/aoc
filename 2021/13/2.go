package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	dots := map[image.Point]struct{}{}
	for _, s := range strings.Fields(split[0]) {
		var p image.Point
		fmt.Sscanf(s, "%d,%d", &p.X, &p.Y)
		dots[p] = struct{}{}
	}

	size := map[rune]int{}
	for i, s := range strings.Split(split[1], "\n") {
		var axis rune
		var line int
		fmt.Sscanf(s, "fold along %c=%d", &axis, &line)

		for p := range dots {
			coord := map[rune]*int{'x': &p.X, 'y': &p.Y}[axis]
			if *coord < line {
				continue
			}
			delete(dots, p)
			*coord = 2*line - *coord
			dots[p] = struct{}{}
			size[axis] = line
		}

		if i == 0 {
			fmt.Println(len(dots))
		}
	}

	for y := 0; y < size['y']; y++ {
		for x := 0; x < size['x']; x++ {
			_, ok := dots[image.Point{x, y}]
			fmt.Print(map[bool]string{true: "██", false: "  "}[ok])
		}
		fmt.Println()
	}
}

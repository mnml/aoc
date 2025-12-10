package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	points, rects, lines := []image.Point{}, []image.Rectangle{}, []image.Rectangle{}
	for i, s := range strings.Fields(string(input)) {
		var p image.Point
		fmt.Sscanf(s, "%d,%d", &p.X, &p.Y)

		for _, q := range points {
			rects = append(rects, image.Rectangle{p, q}.Canon())
		}
		points = append(points, p)
		if len(points) > 1 {
			lines = append(lines, image.Rectangle{points[i-1], points[i]}.Canon())
		}
	}
	lines = append(lines, image.Rectangle{points[len(points)-1], points[0]}.Canon())

	part1, part2 := 0, 0
loop:
	for _, r := range rects {
		r.Max = r.Max.Add(image.Point{1, 1})
		area := r.Dx() * r.Dy()
		part1 = max(part1, area)

		for _, l := range lines {
			l.Max = l.Max.Add(image.Point{1, 1})
			if l.Overlaps(r.Inset(1)) {
				continue loop
			}
		}
		part2 = max(part2, area)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

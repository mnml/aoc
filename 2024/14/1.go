package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type Robot struct {
	P, V image.Point
}

func main() {
	input, _ := os.ReadFile("input.txt")
	area := image.Rectangle{image.Point{0, 0}, image.Point{101, 103}}

	robots, quads := []Robot{}, map[image.Point]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var r Robot
		fmt.Sscanf(s, "p=%d,%d v=%d,%d", &r.P.X, &r.P.Y, &r.V.X, &r.V.Y)
		robots = append(robots, r)
		r.P = r.P.Add(r.V.Mul(100)).Mod(area)
		quads[image.Point{sgn(r.P.X - area.Dx()/2), sgn(r.P.Y - area.Dy()/2)}]++
	}
	fmt.Println(quads[image.Point{-1, -1}] * quads[image.Point{1, -1}] *
		quads[image.Point{1, 1}] * quads[image.Point{-1, 1}])

	for t := 1; ; t++ {
		seen := map[image.Point]struct{}{}
		for i := range robots {
			robots[i].P = robots[i].P.Add(robots[i].V).Mod(area)
			seen[robots[i].P] = struct{}{}
		}
		if len(seen) == len(robots) {
			fmt.Println(t)
			break
		}
	}
}

func sgn(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}

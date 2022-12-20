package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	jets := strings.TrimSpace(string(input))
	rocks := [][]image.Point{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}},
		{{2, 2}, {2, 1}, {0, 0}, {1, 0}, {2, 0}},
		{{0, 3}, {0, 2}, {0, 1}, {0, 0}},
		{{0, 1}, {1, 1}, {0, 0}, {1, 0}},
	}

	grid := map[image.Point]struct{}{}
	move := func(rock []image.Point, delta image.Point) bool {
		nrock := make([]image.Point, len(rock))
		for i, p := range rock {
			p = p.Add(delta)
			if _, ok := grid[p]; ok || p.X < 0 || p.X >= 7 || p.Y < 0 {
				return false
			}
			nrock[i] = p
		}
		copy(rock, nrock)
		return true
	}

	cache := map[[2]int][]int{}

	height, jet := 0, 0
	for i := 0; i < 1000000000000; i++ {
		if i == 2022 {
			fmt.Println(height)
		}

		k := [2]int{i % len(rocks), jet}
		if c, ok := cache[k]; ok {
			if n, d := 1000000000000-i, i-c[0]; n%d == 0 {
				fmt.Println(height + n/d*(height-c[1]))
				break
			}
		}
		cache[k] = []int{i, height}

		rock := []image.Point{}
		for _, p := range rocks[i%len(rocks)] {
			rock = append(rock, p.Add(image.Point{2, height + 3}))
		}

		for {
			move(rock, image.Point{int(jets[jet]) - int('='), 0})
			jet = (jet + 1) % len(jets)

			if !move(rock, image.Point{0, -1}) {
				for _, p := range rock {
					grid[p] = struct{}{}
					if p.Y+1 > height {
						height = p.Y + 1
					}
				}
				break
			}
		}
	}
}

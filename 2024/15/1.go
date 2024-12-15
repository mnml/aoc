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
	r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")

	fmt.Println(run(split[0], split[1]))
	fmt.Println(run(r.Replace(split[0]), split[1]))
}

func run(input, moves string) int {
	grid, robot := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			if r == '@' {
				robot = image.Point{x, y}
				r = '.'
			}
			grid[image.Point{x, y}] = r
		}
	}

	delta := map[rune]image.Point{
		'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		'[': {1, 0}, ']': {-1, 0},
	}

loop:
	for _, r := range strings.ReplaceAll(moves, "\n", "") {
		queue, boxes := []image.Point{robot}, map[image.Point]rune{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if _, ok := boxes[p]; ok {
				continue
			}
			boxes[p] = grid[p]

			switch n := p.Add(delta[r]); grid[n] {
			case '#':
				continue loop
			case '[', ']':
				queue = append(queue, n.Add(delta[grid[n]]))
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for b := range boxes {
			grid[b] = '.'
		}
		for b := range boxes {
			grid[b.Add(delta[r])] = boxes[b]
		}
		robot = robot.Add(delta[r])
	}

	gps := 0
	for p, r := range grid {
		if r == 'O' || r == '[' {
			gps += 100*p.Y + p.X
		}
	}
	return gps
}

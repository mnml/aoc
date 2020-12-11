package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	seats := map[image.Point]rune{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			seats[image.Point{x, y}] = r
		}
	}

	fmt.Println(run(seats, 4, func(p, d image.Point) image.Point { return p.Add(d) }))
	fmt.Println(run(seats, 5, func(p, d image.Point) image.Point {
		for seats[p.Add(d)] == '.' {
			p = p.Add(d)
		}
		return p.Add(d)
	}))
}

func run(seats map[image.Point]rune, maxAdj int, adj func(p, d image.Point) image.Point) (occ int) {
	delta := []image.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for diff := true; diff; {
		occ, diff = 0, false

		next := map[image.Point]rune{}
		for p, r := range seats {
			sum := 0
			for _, d := range delta {
				if seats[adj(p, d)] == '#' {
					sum++
				}
			}

			if r == '#' && sum >= maxAdj {
				r = 'L'
			} else if r == 'L' && sum == 0 || r == '#' {
				r = '#'
				occ++
			}
			next[p] = r
			diff = diff || next[p] != seats[p]
		}
		seats = next
	}
	return
}

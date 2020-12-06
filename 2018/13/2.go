package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"sort"
	"strings"
)

type Cart struct {
	Dir   int
	State int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	grid := strings.Split(strings.TrimRight(string(input), "\n"), "\n")

	carts := map[image.Point]Cart{}
	for y, s := range grid {
		for x, r := range s {
			if p, ok := map[rune]int{'^': 0, '>': 1, 'v': 2, '<': 3}[r]; ok {
				carts[image.Point{x, y}] = Cart{Dir: p}
			}
		}
	}

	var firstCrash *image.Point
	for len(carts) > 1 {
		ks := []image.Point{}
		for p := range carts {
			ks = append(ks, p)
		}
		sort.Slice(ks, func(i, j int) bool {
			return ks[i].Y < ks[j].Y || ks[i].Y == ks[j].Y && ks[i].X < ks[j].X
		})

		for _, p := range ks {
			nc, ok := carts[p]
			if !ok {
				continue
			}
			delete(carts, p)

			np := p.Add([]image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}[nc.Dir])
			if _, ok := carts[np]; ok {
				if firstCrash == nil {
					firstCrash = &np
				}
				delete(carts, np)
				continue
			}

			switch grid[np.Y][np.X] {
			case '/':
				nc.Dir = []int{1, 0, 3, 2}[nc.Dir]
			case '\\':
				nc.Dir = []int{3, 2, 1, 0}[nc.Dir]
			case '+':
				nc.Dir = (nc.Dir + 3 + nc.State) % 4
				nc.State = (nc.State + 1) % 3
			}
			carts[np] = nc
		}
	}

	fmt.Printf("%d,%d\n", firstCrash.X, firstCrash.Y)
	for p := range carts {
		fmt.Printf("%d,%d\n", p.X, p.Y)
	}
}

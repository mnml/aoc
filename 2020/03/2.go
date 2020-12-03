package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	grid := map[image.Point]rune{}
	w, h := 0, 0
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
		w, h = len(s), y+1
	}

	slopes := []image.Point{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	coords, ntrees := make([]image.Point, len(slopes)), map[image.Point]int{}
	for y := 0; y < h; y++ {
		for i, s := range slopes {
			if grid[coords[i]] == '#' {
				ntrees[s]++
			}
			coords[i] = coords[i].Add(s)
			coords[i].X %= w
		}
	}

	mult := 1
	for _, v := range ntrees {
		mult *= v
	}
	fmt.Println(ntrees[image.Point{3, 1}])
	fmt.Println(mult)
}

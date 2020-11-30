package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	nclaims := map[image.Point]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var id, x, y, w, h int
		fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				nclaims[image.Point{i, j}]++
			}
		}
	}

	area := 0
	for _, v := range nclaims {
		if v > 1 {
			area++
		}
	}
	fmt.Println(area)

out:
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var id, x, y, w, h int
		fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				if nclaims[image.Point{i, j}] > 1 {
					continue out
				}
			}
		}
		fmt.Println(id)
	}
}

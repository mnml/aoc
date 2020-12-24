package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	delta := map[string]image.Point{
		"e": {2, 0}, "se": {1, 1}, "sw": {-1, 1},
		"w": {-2, 0}, "nw": {-1, -1}, "ne": {1, -1},
	}

	black := map[image.Point]struct{}{}
	for _, s := range strings.Fields(string(input)) {
		p := image.Point{0, 0}
		for _, s := range regexp.MustCompile(`(e|se|sw|w|nw|ne)`).FindAllString(s, -1) {
			p = p.Add(delta[s])
		}

		if _, ok := black[p]; ok {
			delete(black, p)
		} else {
			black[p] = struct{}{}
		}
	}
	fmt.Println(len(black))

	for i := 0; i < 100; i++ {
		neigh := map[image.Point]int{}
		for p := range black {
			for _, d := range delta {
				neigh[p.Add(d)]++
			}
		}

		new := map[image.Point]struct{}{}
		for p, n := range neigh {
			if _, ok := black[p]; ok && n == 1 || n == 2 {
				new[p] = struct{}{}
			}
		}
		black = new
	}
	fmt.Println(len(black))
}

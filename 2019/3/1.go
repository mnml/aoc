package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	wires := strings.Fields(string(input))
	seen := make([]map[image.Point]struct{}, len(wires))
	min := 0

	for i, w := range wires {
		x, y := 0, 0
		seen[i] = map[image.Point]struct{}{}

		for _, s := range strings.Split(w, ",") {
			for j, _ := strconv.Atoi(s[1:]); j > 0; j-- {
				d := map[byte]image.Point{'U': {0, -1}, 'D': {0, 1}, 'L': {-1, 0}, 'R': {1, 0}}[s[0]]
				x, y = x+d.X, y+d.Y
				seen[i][image.Point{x, y}] = struct{}{}
			}
		}
	}

	for p := range seen[1] {
		if _, ok := seen[0][p]; ok {
			dist := int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
			if min == 0 || dist < min {
				min = dist
			}
		}
	}

	fmt.Println(min)
}

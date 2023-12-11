package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Fields(string(input))

	dists := func(expand int) (d int) {
		galaxies := []image.Point{}
		dy := 0
		for y, s := range split {
			if !strings.Contains(s, "#") {
				dy += expand - 1
			}

			dx := 0
			for x, r := range s {
				col := ""
				for _, s := range split {
					col += string(s[x])
				}
				if !strings.Contains(col, "#") {
					dx += expand - 1
				}

				if r == '#' {
					for _, g := range galaxies {
						d += int(math.Abs(float64(x+dx-g.X)) + math.Abs(float64(y+dy-g.Y)))
					}
					galaxies = append(galaxies, image.Point{x + dx, y + dy})
				}
			}
		}
		return
	}

	fmt.Println(dists(2))
	fmt.Println(dists(1000000))
}

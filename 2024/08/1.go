package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	bounds, freq := map[image.Point]bool{}, map[rune][]image.Point{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			bounds[image.Point{x, y}] = true
			if r != '.' {
				freq[r] = append(freq[r], image.Point{x, y})
			}
		}
	}

	part1, part2 := map[image.Point]struct{}{}, map[image.Point]struct{}{}
	for _, antennas := range freq {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a2 == a1 {
					continue
				}
				if a := a2.Add(a2.Sub(a1)); bounds[a] {
					part1[a] = struct{}{}
				}
				for d := a2.Sub(a1); bounds[a2]; a2 = a2.Add(d) {
					part2[a2] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(part1))
	fmt.Println(len(part2))
}

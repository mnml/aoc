package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	grid := strings.Fields(string(input))
	slopes := map[image.Point]int{{1, 1}: 0, {3, 1}: 0, {5, 1}: 0, {7, 1}: 0, {1, 2}: 0}
	mult := 1
	for s := range slopes {
		for p := (image.Point{}); p.Y < len(grid); p = p.Add(s) {
			if grid[p.Y][p.X%len(grid[0])] == '#' {
				slopes[s]++
			}
		}
		mult *= slopes[s]
	}
	fmt.Println(slopes[image.Point{3, 1}])
	fmt.Println(mult)
}

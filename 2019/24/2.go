package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	init := [5][5]int{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '#' {
				init[y][x] = 1
			}
		}
	}
	grid := map[int][5][5]int{0: init}

	var bugs int
	for i := 0; i < 200; i++ {
		bugs = 0

		next := map[int][5][5]int{}
		for z := -200; z <= 200; z++ {
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if y == 2 && x == 2 {
						continue
					}

					adj := 0
					for _, d := range []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
						n := image.Point{x, y}.Add(d)
						switch {
						case n.Y == -1 || n.Y == 5 || n.X == -1 || n.X == 5:
							adj += grid[z-1][2+d.Y][2+d.X]
						case n.Y == 2 && n.X == 2:
							for i := 0; i < 5; i++ {
								adj += grid[z+1][d.X&1*i-d.Y&1*2*(d.Y-1)][d.Y&1*i-d.X&1*2*(d.X-1)]
							}
						default:
							adj += grid[z][n.Y][n.X]
						}
					}

					if adj == 1 || grid[z][y][x] == 0 && adj == 2 {
						temp := next[z]
						temp[y][x] = 1
						next[z] = temp
						bugs++
					}
				}
			}
		}
		grid = next
	}

	fmt.Println(bugs)
}

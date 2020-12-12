package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	ship1, dir := image.Point{0, 0}, 1
	ship2, wp := image.Point{0, 0}, image.Point{10, -1}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var action rune
		var value int
		fmt.Sscanf(s, "%c%d", &action, &value)

		delta := map[rune]image.Point{'N': {0, -1}, 'E': {1, 0}, 'S': {0, 1}, 'W': {-1, 0}}
		switch action {
		case 'N', 'S', 'E', 'W':
			ship1 = ship1.Add(delta[action].Mul(value))
			wp = wp.Add(delta[action].Mul(value))
		case 'L':
			value *= 3
			fallthrough
		case 'R':
			dir = (dir + value/90) % 4
			for i := 0; i < value/90; i++ {
				wp.X, wp.Y = -wp.Y, wp.X
			}
		case 'F':
			ship1 = ship1.Add(delta[rune("NESW"[dir])].Mul(value))
			ship2 = ship2.Add(wp.Mul(value))
		}
	}
	fmt.Println(abs(ship1.X) + abs(ship1.Y))
	fmt.Println(abs(ship2.X) + abs(ship2.Y))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

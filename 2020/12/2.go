package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	ins := strings.Split(strings.TrimSpace(string(input)), "\n")

	ship, wp := image.Point{0, 0}, image.Point{10, -1}
	fmt.Println(run(ins, &ship, &image.Point{1, 0}, &ship))
	fmt.Println(run(ins, &image.Point{0, 0}, &wp, &wp))
}

func run(ins []string, ship, dir, mov *image.Point) int {
	delta := map[rune]image.Point{'N': {0, -1}, 'S': {0, 1}, 'E': {1, 0}, 'W': {-1, 0}, 'L': {-1, 1}, 'R': {1, -1}}

	for _, s := range ins {
		var action rune
		var value int
		fmt.Sscanf(s, "%c%d", &action, &value)

		switch action {
		case 'N', 'S', 'E', 'W':
			*mov = mov.Add(delta[action].Mul(value))
		case 'L', 'R':
			for i := 0; i < value/90; i++ {
				dir.X, dir.Y = delta[action].Y*dir.Y, delta[action].X*dir.X
			}
		case 'F':
			*ship = ship.Add(dir.Mul(value))
		}
	}

	return int(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))
}

package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	delta := map[string]image.Point{"forward": {1, 0}, "down": {0, 1}, "up": {0, -1}}

	part1, part2 := image.Point{0, 0}, image.Point{0, 0}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		cmd, arg := "", 0
		fmt.Sscanf(s, "%s %d", &cmd, &arg)
		part1 = part1.Add(delta[cmd].Mul(arg))
		part2 = part2.Add(image.Point{arg, arg * part1.Y}.Mul(delta[cmd].X))
	}
	fmt.Println(part1.X * part1.Y)
	fmt.Println(part2.X * part2.Y)
}

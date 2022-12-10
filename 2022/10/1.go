package main

import (
	"fmt"
	"os"
	"strings"
)

const w = 40

func main() {
	input, _ := os.ReadFile("input.txt")

	c, x := 0, 1
	part1, part2 := 0, ""

	tick := func() {
		part2 += map[bool]string{true: "#", false: "."}[c%w >= x-1 && c%w <= x+1]
		part2 += map[bool]string{true: "\n"}[c%w == w-1]
		if c++; (c+w/2)%w == 0 {
			part1 += c * x
		}
	}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var ins string
		var v int
		fmt.Sscanf(s, "%s %d", &ins, &v)

		tick()
		if ins == "addx" {
			tick()
			x += v
		}
	}
	fmt.Println(part1)
	fmt.Println(strings.TrimSpace(part2))
}

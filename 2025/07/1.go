package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Fields(string(input))

	beams := make([]int, len(split[0]))
	beams[strings.Index(split[0], "S")] = 1

	part1, part2 := 0, 1
	for _, s := range split {
		for i, r := range s {
			if r == '^' {
				part1 += min(beams[i], 1)
				part2 += beams[i]
				beams[i-1] += beams[i]
				beams[i+1] += beams[i]
				beams[i] = 0
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

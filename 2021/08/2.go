package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	digits := map[int]int{42: 0, 17: 1, 34: 2, 39: 3, 30: 4, 37: 5, 41: 6, 25: 7, 49: 8, 45: 9}

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		split := strings.Split(s, " | ")

		out := 0
		for _, s := range strings.Fields(split[1]) {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				part1++
			}

			hash := 0
			for _, r := range s {
				hash += strings.Count(split[0], string(r))
			}
			out = 10*out + digits[hash]
		}
		part2 += out
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

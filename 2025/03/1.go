package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Fields(string(input)) {
		part1 += jolts(s, 2)
		part2 += jolts(s, 12)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func jolts(bank string, nbat int) int {
	s, p := "", 0
	for b := nbat; b > 0; b-- {
		for i := p; i <= len(bank)-b; i++ {
			if bank[i] > bank[p] {
				p = i
			}
		}
		s += string(bank[p])
		p++
	}
	n, _ := strconv.Atoi(s)
	return n
}

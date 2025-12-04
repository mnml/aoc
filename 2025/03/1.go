package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Fields(string(input)) {
		part1 += joltage(s, 2)
		part2 += joltage(s, 12)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func joltage(bank string, nbat int) int {
	s := ""
	for i, j := 0, len(bank)-nbat+1; j <= len(bank); i, j = i+1, j+1 {
		bs := []byte(bank[i:j])
		i += slices.Index(bs, slices.Max(bs))
		s += string(bank[i])
	}
	n, _ := strconv.Atoi(s)
	return n
}

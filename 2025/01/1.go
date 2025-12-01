package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	delta := map[byte]int{'L': -1, 'R': 1}
	dial := 50

	part1, part2 := 0, 0
	for _, s := range strings.Fields(string(input)) {
		n, _ := strconv.Atoi(s[1:])

		for range n {
			if dial += delta[s[0]]; dial%100 == 0 {
				part2++
			}
		}
		if dial%100 == 0 {
			part1++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

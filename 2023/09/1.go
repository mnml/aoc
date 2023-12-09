package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	var next func([]int) int
	next = func(s []int) int {
		h := []int{}
		for i := 0; i < len(s)-1; i++ {
			h = append(h, s[i+1]-s[i])
		}
		if !slices.ContainsFunc(h, func(x int) bool { return x != 0 }) {
			return s[len(s)-1]
		}
		return s[len(s)-1] + next(h)
	}

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var seq []int
		json.Unmarshal([]byte("["+strings.ReplaceAll(s, " ", ",")+"]"), &seq)

		part1 += next(seq)
		slices.Reverse(seq)
		part2 += next(seq)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

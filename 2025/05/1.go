package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	ranges := [][2]int{}
	for _, s := range strings.Fields(split[0]) {
		var a, b int
		fmt.Sscanf(s, "%d-%d", &a, &b)
		ranges = append(ranges, [2]int{a, b})
	}
	slices.SortFunc(ranges, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })

	part1 := 0
	for _, s := range strings.Fields(split[1]) {
		n, _ := strconv.Atoi(s)
		for _, r := range ranges {
			if r[0] <= n && n <= r[1] {
				part1++
				break
			}
		}
	}
	fmt.Println(part1)

	part2, c := 0, 0
	for _, r := range ranges {
		part2 += max(c, r[1]+1) - max(c, r[0])
		c = max(c, r[1]+1)
	}
	fmt.Println(part2)
}

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

	comp := [][2]int{ranges[0]}
	for _, r := range ranges[1:] {
		if last := &comp[len(comp)-1]; r[0] <= last[1] {
			last[1] = max(last[1], r[1])
		} else {
			comp = append(comp, r)
		}
	}

	part2 := 0
	for _, r := range comp {
		part2 += r[1] - r[0] + 1
	}
	fmt.Println(part2)
}

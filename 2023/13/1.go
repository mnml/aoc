package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		rows, cols := []string{}, make([]string, len(strings.Fields(s)[0]))
		for _, s := range strings.Fields(s) {
			rows = append(rows, s)
			for i, r := range s {
				cols[i] += string(r)
			}
		}

		part1 += mirror(cols, slices.Equal) + 100*mirror(rows, slices.Equal)
		part2 += mirror(cols, smudge) + 100*mirror(rows, smudge)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func mirror(s []string, equal func([]string, []string) bool) int {
	for i := 1; i < len(s); i++ {
		l := slices.Min([]int{i, len(s) - i})
		a, b := slices.Clone(s[i-l:i]), s[i:i+l]
		slices.Reverse(a)
		if equal(a, b) {
			return i
		}
	}
	return 0
}

func smudge(a, b []string) bool {
	diffs := 0
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				diffs++
			}
		}
	}
	return diffs == 1
}

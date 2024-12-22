package main

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, map[[4]int]int{}
	for _, s := range strings.Fields(string(input)) {
		n, _ := strconv.Atoi(s)

		secrets, diff := []int{n}, []int{}
		for i := range 2000 {
			n ^= n * 64 % 16777216
			n ^= n / 32 % 16777216
			n ^= n * 2048 % 16777216
			secrets = append(secrets, n)
			diff = append(diff, n%10-secrets[i]%10)
		}
		part1 += n

		seen := map[[4]int]bool{}
		for i := range len(secrets) - 4 {
			if p := [4]int(diff[i : i+4]); !seen[p] {
				part2[p] += secrets[i+4] % 10
				seen[p] = true
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2[slices.MaxFunc(slices.Collect(maps.Keys(part2)),
		func(a, b [4]int) int { return cmp.Compare(part2[a], part2[b]) })])
}

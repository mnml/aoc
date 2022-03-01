package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	points := map[byte]int{
		')': 3, ']': 57, '}': 1197, '>': 25137,
		'(': 1, '[': 2, '{': 3, '<': 4,
	}
	re := regexp.MustCompile(`\(\)|\[]|{}|<>`)

	part1, part2 := 0, []int{}
	for _, s := range strings.Fields(string(input)) {
		for re.MatchString(s) {
			s = re.ReplaceAllString(s, "")
		}

		if i := strings.IndexAny(s, ")]}>"); i != -1 {
			part1 += points[s[i]]
			continue
		}

		score := 0
		for i := len(s) - 1; i >= 0; i-- {
			score = score*5 + points[s[i]]
		}
		part2 = append(part2, score)
	}
	fmt.Println(part1)
	sort.Ints(part2)
	fmt.Println(part2[len(part2)/2])
}

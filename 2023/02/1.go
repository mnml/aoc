package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`(\d+) (\w+)`)

	part1, part2 := 0, 0
	for i, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		mins := map[string]int{}

		for _, m := range re.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(m[1])
			mins[m[2]] = max(mins[m[2]], n)
		}

		if mins["red"] <= 12 && mins["green"] <= 13 && mins["blue"] <= 14 {
			part1 += i + 1
		}
		part2 += mins["red"] * mins["green"] * mins["blue"]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

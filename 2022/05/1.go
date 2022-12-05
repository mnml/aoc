package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(string(input), "\n\n")
	crates := strings.Split(split[0], "\n")
	keys := crates[len(crates)-1]

	stack1, stack2 := map[rune]string{}, map[rune]string{}
	for _, s := range crates {
		for i, r := range s {
			if unicode.IsLetter(r) {
				stack1[[]rune(keys)[i]] += string(r)
				stack2[[]rune(keys)[i]] += string(r)
			}
		}
	}

	for _, s := range strings.Split(strings.TrimSpace(split[1]), "\n") {
		var qty int
		var from, to rune
		fmt.Sscanf(s, "move %d from %c to %c", &qty, &from, &to)

		for i := 0; i < qty; i++ {
			stack1[to] = stack1[from][:1] + stack1[to]
			stack1[from] = stack1[from][1:]
		}
		stack2[to] = stack2[from][:qty] + stack2[to]
		stack2[from] = stack2[from][qty:]
	}

	part1, part2 := "", ""
	for _, r := range strings.ReplaceAll(keys, " ", "") {
		part1 += stack1[r][:1]
		part2 += stack2[r][:1]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

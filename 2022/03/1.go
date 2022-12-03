package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Fields(strings.TrimSpace(string(input)))

	part1, part2 := 0, 0
	for i, s := range split {
		part1 += common(s[:len(s)/2], s[len(s)/2:])

		if i%3 == 0 {
			part2 += common(split[i : i+3]...)
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func common(strs ...string) int {
loop:
	for _, r := range strs[0] {
		for _, s := range strs[1:] {
			if !strings.ContainsRune(s, r) {
				continue loop
			}
		}
		return strings.IndexRune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r)
	}
	return 0
}

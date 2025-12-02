package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		var a, b int
		fmt.Sscanf(s, "%d-%d", &a, &b)

		for i := a; i <= b; i++ {
			s := strconv.Itoa(i)
			if s[:len(s)/2] == s[len(s)/2:] {
				part1 += i
			}
			if strings.Contains((s + s)[1:len(s+s)-1], s) {
				part2 += i
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

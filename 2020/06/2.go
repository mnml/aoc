package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		yes := map[rune]int{}
		for _, s := range strings.Fields(s) {
			for _, r := range s {
				yes[r]++
			}
		}
		part1 += len(yes)

		for _, v := range yes {
			if v == len(strings.Fields(s)) {
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

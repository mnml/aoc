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
		people := strings.Fields(s)
		for _, s := range people {
			for _, r := range s {
				yes[r]++
				if yes[r] == len(people) {
					part2++
				}
			}
		}
		part1 += len(yes)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

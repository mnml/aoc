package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	boxes := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		count := map[rune]int{}
		for _, r := range s {
			count[r]++
		}
		has := map[int]struct{}{}
		for _, v := range count {
			if _, ok := has[v]; !ok {
				boxes[v] += 1
			}
			has[v] = struct{}{}
		}
	}
	fmt.Println(boxes[2] * boxes[3])
}

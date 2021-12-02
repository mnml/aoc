package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	depths := []int{}
	for _, s := range strings.Fields(string(input)) {
		i, _ := strconv.Atoi(s)
		depths = append(depths, i)
	}

	part1, part2 := 0, 0
	for i := range depths {
		if i >= 1 && depths[i] > depths[i-1] {
			part1++
		}
		if i >= 3 && depths[i] > depths[i-3] {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

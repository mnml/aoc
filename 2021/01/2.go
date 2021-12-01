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
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			part1++
		}

		if i < len(depths)-2 &&
			depths[i]+depths[i+1]+depths[i+2] > depths[i-1]+depths[i]+depths[i+1] {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

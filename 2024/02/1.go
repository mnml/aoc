package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		report := []int{}
		for _, s := range strings.Fields(s) {
			n, _ := strconv.Atoi(s)
			report = append(report, n)
		}

		if check(report) {
			part1++
		}
		for i := range report {
			if check(append(slices.Clone(report[:i]), report[i+1:]...)) {
				part2++
				break
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func check(r []int) bool {
	for i := 1; i < len(r); i++ {
		d := r[i] - r[i-1]
		if sgn(d) != sgn(r[1]-r[0]) || abs(d) < 1 || abs(d) > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

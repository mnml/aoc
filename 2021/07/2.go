package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	crabs, max := []int{}, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		p, _ := strconv.Atoi(s)
		crabs = append(crabs, p)
		if p > max {
			max = p
		}
	}
	fmt.Println(run(crabs, max, func(p int) int { return p }))
	fmt.Println(run(crabs, max, func(p int) int { return (p * (p + 1)) / 2 }))
}

func run(pos []int, max int, f func(int) int) int {
	min := int(^uint(0) >> 1)
	for b := 0; b <= max; b++ {
		sum := 0
		for _, a := range pos {
			sum += f(abs(b - a))
		}
		if sum < min {
			min = sum
		}
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

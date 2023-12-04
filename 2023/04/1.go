package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`((?:\s+\d+)*) \|((?:\s+\d+)*)`)
	copies := map[int]int{}

	part1, part2 := 0, 0
	for i, m := range re.FindAllStringSubmatch(string(input), -1) {
		matching := 0
		for _, n := range strings.Fields(m[1]) {
			if slices.Contains(strings.Fields(m[2]), n) {
				matching++
			}
		}
		part1 += int(math.Pow(2, float64(matching-1)))

		copies[i]++
		part2++
		for j := 1; j <= matching; j++ {
			copies[i+j] += copies[i]
			part2 += copies[i]
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

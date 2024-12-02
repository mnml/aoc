package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var report []int
		json.Unmarshal([]byte("["+strings.ReplaceAll(s, " ", ",")+"]"), &report)

		if check(report) {
			part1++
		}
		for i := range report {
			if check(slices.Delete(slices.Clone(report), i, i+1)) {
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
		if d := r[i] - r[i-1]; d*(r[1]-r[0]) <= 0 || d < -3 || d > 3 {
			return false
		}
	}
	return true
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	sum := 0
	for _, s := range strings.Fields(string(input)) {
		n := 0
		for _, r := range s {
			n = 5*n + map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}[r]
		}
		sum += n
	}

	snafu := ""
	for sum > 0 {
		snafu = string("=-012"[(sum+2)%5]) + snafu
		sum = (sum + 2) / 5
	}
	fmt.Println(snafu)
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	cals := make([]int, len(split))
	for i, s := range split {
		for _, s := range strings.Fields(s) {
			c, _ := strconv.Atoi(s)
			cals[i] += c
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))
	fmt.Println(cals[0])
	fmt.Println(cals[0] + cals[1] + cals[2])
}

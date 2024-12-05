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
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	cmp := func(a, b string) int {
		for _, s := range strings.Split(split[0], "\n") {
			if s := strings.Split(s, "|"); s[0] == a && s[1] == b {
				return -1
			}
		}
		return 0
	}

	run := func(sorted bool) (r int) {
		for _, s := range strings.Split(split[1], "\n") {
			if s := strings.Split(s, ","); slices.IsSortedFunc(s, cmp) == sorted {
				slices.SortFunc(s, cmp)
				n, _ := strconv.Atoi(s[len(s)/2])
				r += n
			}
		}
		return r
	}

	fmt.Println(run(true))
	fmt.Println(run(false))
}

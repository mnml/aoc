package main

import (
	"encoding/json"
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
		s := strings.Split(s, ": ")
		test, _ := strconv.Atoi(s[0])
		var numbers []int
		json.Unmarshal([]byte("["+strings.ReplaceAll(s[1], " ", ",")+"]"), &numbers)

		for _, op := range perm(3, len(numbers)-1) {
			result := numbers[0]
			for i, n := range numbers[1:] {
				switch op[i] {
				case 0:
					result += n
				case 1:
					result *= n
				case 2:
					result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(n))
				}
			}
			if result == test {
				if !slices.Contains(op, 2) {
					part1 += test
				}
				part2 += test
				break
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func perm(n, l int) (r [][]int) {
	if l == 0 {
		return [][]int{{}}
	}
	for _, p := range perm(n, l-1) {
		for i := range n {
			r = append(r, append([]int{i}, p...))
		}
	}
	return r
}

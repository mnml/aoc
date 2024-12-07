package main

import (
	"encoding/json"
	"fmt"
	"os"
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
		part1 += value(test, numbers, false)
		part2 += value(test, numbers, true)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func value(test int, ns []int, p2 bool) int {
	if len(ns) == 1 {
		if ns[0] == test {
			return test
		}
		return 0
	}
	if n := value(test, append([]int{ns[0] + ns[1]}, ns[2:]...), p2); n != 0 {
		return n
	}
	if n := value(test, append([]int{ns[0] * ns[1]}, ns[2:]...), p2); n != 0 {
		return n
	}
	if n, _ := strconv.Atoi(fmt.Sprintf("%d%d", ns[0], ns[1])); p2 {
		return value(test, append([]int{n}, ns[2:]...), p2)
	}
	return 0
}

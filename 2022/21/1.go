package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var monkeys = map[string]string{}

func main() {
	input, _ := os.ReadFile("input.txt")

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		s := strings.Split(s, ": ")
		monkeys[s[0]] = s[1]
	}
	fmt.Println(solve("root"))

	monkeys["humn"] = "0"
	s := strings.Fields(monkeys["root"])
	if solve(s[0]) < solve(s[2]) {
		s[0], s[2] = s[2], s[0]
	}

	part2, _ := sort.Find(1e16, func(v int) int {
		monkeys["humn"] = strconv.Itoa(v)
		return solve(s[0]) - solve(s[2])
	})
	fmt.Println(part2)
}

func solve(expr string) int {
	if v, err := strconv.Atoi(monkeys[expr]); err == nil {
		return v
	}

	s := strings.Fields(monkeys[expr])
	return map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}[s[1]](solve(s[0]), solve(s[2]))
}

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		part1 += run(s, regexp.MustCompile(`\([^\(\)]+\)`), eval)
		part2 += run(s, regexp.MustCompile(`\([^\(\)]+\)`), func(s string) int {
			return run(s, regexp.MustCompile(`\d+ \+ \d+`), eval)
		})
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func run(s string, re *regexp.Regexp, eval func(string) int) int {
	for re.MatchString(s) {
		s = re.ReplaceAllStringFunc(s, func(s string) string {
			return strconv.Itoa(eval(s))
		})
	}
	return eval(s)
}

func eval(s string) int {
	fields := strings.Fields(strings.Trim(s, "()"))
	acc, _ := strconv.Atoi(fields[0])

	for i := 1; i < len(fields); i += 2 {
		switch n, _ := strconv.Atoi(fields[i+1]); fields[i] {
		case "+":
			acc += n
		case "*":
			acc *= n
		}
	}
	return acc
}

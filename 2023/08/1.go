package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

	network := map[string]map[rune]string{}
	for _, m := range re.FindAllStringSubmatch(split[1], -1) {
		network[m[1]] = map[rune]string{'L': m[2], 'R': m[3]}
	}

	walk := func(start, end string) int {
		result := 1
		for n := range network {
			if !strings.HasSuffix(n, start) {
				continue
			}

			steps := 0
			for !strings.HasSuffix(n, end) {
				n = network[n][rune(split[0][steps%len(split[0])])]
				steps++
			}

			result = lcm(result, steps)
		}
		return result
	}

	fmt.Println(walk("AAA", "ZZZ"))
	fmt.Println(walk("A", "Z"))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

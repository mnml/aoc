package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`(?s)(?:^|do\(\)).*?(?:don't\(\)|$)`)

	part2 := 0
	for _, m := range re.FindAllString(string(input), -1) {
		part2 += mul(m)
	}
	fmt.Println(mul(string(input)))
	fmt.Println(part2)
}

func mul(s string) (r int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, m := range re.FindAllStringSubmatch(s, -1) {
		n1, _ := strconv.Atoi(m[1])
		n2, _ := strconv.Atoi(m[2])
		r += n1 * n2
	}
	return r
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`(?s)(?:^|do\(\)).*?(?:don't\(\)|$)`)
	fmt.Println(mul(string(input)))
	fmt.Println(mul(strings.Join(re.FindAllString(string(input), -1), "")))
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

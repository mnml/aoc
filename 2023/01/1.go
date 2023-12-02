package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	calc := func(r *strings.Replacer) (result int) {
		for _, s := range strings.Fields(string(input)) {
			s = r.Replace(r.Replace(s))
			result += 10 * int(s[strings.IndexAny(s, "123456789")]-'0')
			result += int(s[strings.LastIndexAny(s, "123456789")] - '0')
		}
		return
	}

	fmt.Println(calc(strings.NewReplacer()))
	fmt.Println(calc(strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")))
}

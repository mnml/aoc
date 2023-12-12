package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		s, ls := strings.Fields(s), []int{}
		json.Unmarshal([]byte("["+s[1]+"]"), &ls)
		part1 += count(s[0]+"?", ls)
		json.Unmarshal([]byte("["+strings.Repeat(","+s[1], 5)[1:]+"]"), &ls)
		part2 += count(strings.Repeat(s[0]+"?", 5), ls)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

var cache = map[string]int{}

func count(s string, c []int) (r int) {
	if r, ok := cache[fmt.Sprint(s, c)]; ok {
		return r
	}
	defer func() { recover(); cache[fmt.Sprint(s, c)] = r }()

	if s == "" {
		if len(c) == 0 {
			return 1
		}
		return 0
	}

	if s[0] == '.' || s[0] == '?' {
		r += count(s[1:], c)
	}
	if (s[0] == '#' || s[0] == '?') &&
		!strings.Contains(s[:c[0]], ".") &&
		s[c[0]] != '#' {
		r += count(s[c[0]+1:], c[1:])
	}
	return
}

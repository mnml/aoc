package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Cat  string
	Op   string
	Val  int
	Next string
}

func main() {
	input, _ := os.ReadFile("input.txt")
	re1 := regexp.MustCompile(`(\w+){(.*),(\w+)}`)
	re2 := regexp.MustCompile(`(\w+)(<|>)(\d+):(\w+)`)

	workflows := map[string][]Rule{}
	for _, m := range re1.FindAllStringSubmatch(string(input), -1) {
		w := m[1]
		for _, m := range re2.FindAllStringSubmatch(m[2], -1) {
			v, _ := strconv.Atoi(m[3])
			workflows[w] = append(workflows[w], Rule{m[1], m[2], v, m[4]})
		}
		workflows[w] = append(workflows[w], Rule{Next: m[3]})
	}

	var comb func(string, map[string][2]int) int
	comb = func(workflow string, ranges map[string][2]int) (c int) {
		for _, r := range ranges {
			if r[1] < r[0] {
				return 0
			}
		}
		if workflow == "R" {
			return 0
		}
		if workflow == "A" {
			c = 1
			for _, r := range ranges {
				c *= r[1] - r[0] + 1
			}
			return
		}

		for _, r := range workflows[workflow] {
			next := maps.Clone(ranges)

			switch r.Op {
			case "<":
				min := slices.Min([]int{next[r.Cat][1], r.Val - 1})
				max := slices.Max([]int{next[r.Cat][0], r.Val})
				next[r.Cat] = [2]int{next[r.Cat][0], min}
				ranges[r.Cat] = [2]int{max, ranges[r.Cat][1]}
			case ">":
				min := slices.Min([]int{next[r.Cat][1], r.Val})
				max := slices.Max([]int{next[r.Cat][0], r.Val + 1})
				next[r.Cat] = [2]int{max, next[r.Cat][1]}
				ranges[r.Cat] = [2]int{ranges[r.Cat][0], min}
			}
			c += comb(r.Next, next)
		}
		return c
	}

	part1 := 0
	for _, s := range strings.Fields(strings.Split(strings.TrimSpace(string(input)), "\n\n")[1]) {
		s = strings.NewReplacer(`{`, `{"`, `,`, `,"`, `=`, `":`).Replace(s)
		var rating map[string]int
		json.Unmarshal([]byte(s), &rating)

		ranges := map[string][2]int{}
		for c, r := range rating {
			ranges[c] = [2]int{r, r}
		}
		if comb("in", ranges) == 1 {
			part1 += rating["x"] + rating["m"] + rating["a"] + rating["s"]
		}
	}
	fmt.Println(part1)

	ranges := map[string][2]int{
		"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000},
	}
	fmt.Println(comb("in", ranges))
}

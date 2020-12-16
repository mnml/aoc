package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	rules := map[string][]int{}
	for _, s := range strings.Split(split[0], "\n") {
		rule := strings.Split(s, ": ")
		rules[rule[0]] = make([]int, 4)
		fmt.Sscanf(rule[1], "%d-%d or %d-%d", &rules[rule[0]][0], &rules[rule[0]][1], &rules[rule[0]][2], &rules[rule[0]][3])
	}

	indices := map[string]map[int]struct{}{}
	for k := range rules {
		indices[k] = map[int]struct{}{}
		for i := 0; i < len(rules); i++ {
			indices[k][i] = struct{}{}
		}
	}

	part1 := 0
tickets:
	for _, s := range strings.Split(split[2], "\n")[1:] {
	fields:
		for _, s := range strings.Split(s, ",") {
			n, _ := strconv.Atoi(s)
			for _, v := range rules {
				if n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3] {
					continue fields
				}
			}
			part1 += n
			continue tickets
		}

		for i, s := range strings.Split(s, ",") {
			for k, v := range rules {
				if n, _ := strconv.Atoi(s); !(n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3]) {
					delete(indices[k], i)
				}
			}
		}
	}
	fmt.Println(part1)

	part2 := 1
	for len(indices) > 0 {
		for k, v := range indices {
			if len(v) != 1 {
				continue
			}

			for i := range v {
				for k := range indices {
					delete(indices[k], i)
				}
				delete(indices, k)

				if strings.HasPrefix(k, "departure") {
					n, _ := strconv.Atoi(strings.Split(strings.Split(split[1], "\n")[1], ",")[i])
					part2 *= n
				}
			}
		}
	}
	fmt.Println(part2)
}

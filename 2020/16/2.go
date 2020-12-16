package main

import (
	"fmt"
	"io/ioutil"
	"math/bits"
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

	masks := map[string]uint{}
	for k := range rules {
		masks[k] = 1<<len(rules) - 1
	}

	part1 := 0
out:
	for _, s := range strings.Fields(split[2])[2:] {
		invalid := map[string]uint{}

		for i, s := range strings.Split(s, ",") {
			n, _ := strconv.Atoi(s)
			for k, v := range rules {
				if !(n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3]) {
					invalid[k] |= 1 << i
				}
			}

			if len(invalid) == len(rules) {
				part1 += n
				continue out
			}
		}

		for k, v := range invalid {
			masks[k] &^= v
		}
	}
	fmt.Println(part1)

	part2 := 1
	for used := uint(0); used != 1<<len(rules)-1; {
		for k := range masks {
			if masks[k] &^= used; bits.OnesCount(masks[k]) == 1 {
				used |= masks[k]

				if strings.HasPrefix(k, "departure") {
					n, _ := strconv.Atoi(strings.Split(strings.Fields(split[1])[2], ",")[bits.TrailingZeros(masks[k])])
					part2 *= n
				}
			}
		}
	}
	fmt.Println(part2)
}

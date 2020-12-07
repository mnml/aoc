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

	bags := map[string]map[string]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		out := strings.Split(s, " bags ")[0]
		bags[out] = map[string]int{}
		for _, in := range regexp.MustCompile(`(\d+) (\w+ \w+)`).FindAllStringSubmatch(s, -1) {
			bags[out][in[2]], _ = strconv.Atoi(in[1])
		}
	}

	fmt.Println(len(parents(bags, "shiny gold")))
	fmt.Println(count(bags, "shiny gold"))
}

func parents(bags map[string]map[string]int, bag string) map[string]struct{} {
	set := map[string]struct{}{}
	for out := range bags {
		for in := range bags[out] {
			if in == bag {
				set[out] = struct{}{}
				for b := range parents(bags, out) {
					set[b] = struct{}{}
				}
				break
			}
		}
	}
	return set
}

func count(bags map[string]map[string]int, bag string) (total int) {
	for b, c := range bags[bag] {
		total += c * (count(bags, b) + 1)
	}
	return
}

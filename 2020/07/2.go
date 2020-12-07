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

	total := 0
	for b := range bags {
		if contains(bags, b, "shiny gold") {
			total++
		}
	}
	fmt.Println(total)
	fmt.Println(count(bags, "shiny gold"))
}

func contains(bags map[string]map[string]int, out, in string) bool {
	if _, ok := bags[out][in]; ok {
		return true
	}
	for out := range bags[out] {
		if contains(bags, out, in) {
			return true
		}
	}
	return false
}

func count(bags map[string]map[string]int, bag string) (total int) {
	for k, v := range bags[bag] {
		total += v * (count(bags, k) + 1)
	}
	return
}

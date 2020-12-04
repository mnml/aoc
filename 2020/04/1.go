package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	valid := 0
out:
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		for _, f := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if ok := strings.Contains(s, f); !ok {
				continue out
			}
		}
		valid++
	}
	fmt.Println(valid)
}

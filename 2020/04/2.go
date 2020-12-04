package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var res = []*regexp.Regexp{
	regexp.MustCompile(`byr:(19[2-9]\d|200[0-2])(\s+|$)`),
	regexp.MustCompile(`iyr:(201\d|2020)(\s+|$)`),
	regexp.MustCompile(`eyr:(202\d|2030)(\s+|$)`),
	regexp.MustCompile(`hgt:((1[5-8]\d|19[0-3])cm|(59|6\d|7[0-6])in)(\s+|$)`),
	regexp.MustCompile(`hcl:#[\da-f]{6}(\s+|$)`),
	regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)(\s+|$)`),
	regexp.MustCompile(`pid:\d{9}(\s+|$)`),
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	valid := 0
out:
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		for _, re := range res {
			if !re.MatchString(s) {
				continue out
			}
		}
		valid++
	}
	fmt.Println(valid)
}

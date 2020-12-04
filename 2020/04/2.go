package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var res = []*regexp.Regexp{
	regexp.MustCompile(`(?:^|\s)(byr):(?:(19[2-9]\d|200[0-2])(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(iyr):(?:(201\d|2020)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(eyr):(?:(202\d|2030)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(hgt):(?:((?:1[5-8]\d|19[0-3])cm|(?:59|6\d|7[0-6])in)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(hcl):(?:(#[\da-f]{6})(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(ecl):(?:(amb|blu|brn|gry|grn|hzl|oth)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(pid):(?:(\d{9})(?:\s|$))?`),
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		d1, d2 := 1, 1
		for _, re := range res {
			if m := re.FindStringSubmatch(s); len(m) == 0 {
				d1, d2 = 0, 0
			} else if m[2] == "" {
				d2 = 0
			}
		}
		part1, part2 = part1+d1, part2+d2
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

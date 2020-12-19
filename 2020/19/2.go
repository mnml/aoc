package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var rules = map[string]string{}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	for _, s := range strings.Split(strings.TrimSpace(split[0]), "\n") {
		rule := strings.Split(s, ": ")
		rules[rule[0]] = rule[1]
	}

	fmt.Println(len(regexp.MustCompile("(?m)^"+regex("0")+"$").FindAllString(split[1], -1)))

	rules["8"] = `"` + regex("42") + `+"`
	rules["11"] = ""
	for i := 1; i <= 10; i++ {
		rules["11"] += fmt.Sprintf("|%s{%d}%s{%d}", regex("42"), i, regex("31"), i)
	}
	rules["11"] = `"(?:` + rules["11"][1:] + `)"`

	fmt.Println(len(regexp.MustCompile("(?m)^"+regex("0")+"$").FindAllString(split[1], -1)))
}

func regex(rule string) (re string) {
	if rules[rule][0] == '"' {
		return rules[rule][1 : len(rules[rule])-1]
	}
	for _, s := range strings.Split(rules[rule], " | ") {
		re += "|"
		for _, s := range strings.Fields(s) {
			re += regex(s)
		}
	}
	return "(?:" + re[1:] + ")"
}

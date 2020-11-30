package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	steps := map[rune]map[rune]struct{}{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var a, b rune
		fmt.Sscanf(s, "Step %c must be finished before step %c can begin.", &a, &b)
		if _, ok := steps[a]; !ok {
			steps[a] = map[rune]struct{}{}
		}
		if _, ok := steps[b]; !ok {
			steps[b] = map[rune]struct{}{}
		}
		steps[b][a] = struct{}{}
	}

	l := []rune{}
	for len(steps) > 0 {
		n := 'Z'
		for k, v := range steps {
			if k < n && len(v) == 0 {
				n = k
			}
		}
		delete(steps, n)
		l = append(l, n)
		for _, v := range steps {
			delete(v, n)
		}
	}
	fmt.Println(string(l))
}

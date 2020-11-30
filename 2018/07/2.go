package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	workers, seconds := map[rune]int{}, 0
	for len(steps) > 0 || len(workers) > 0 {
		l := []rune{}
		for k, v := range steps {
			if _, ok := workers[k]; !ok && len(v) == 0 {
				l = append(l, k)
			}
		}
		sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })

		for len(l) > 0 && len(workers) < 5 {
			workers[l[0]] = int(l[0]) - 4
			l = l[1:]
		}

		for w := range workers {
			if workers[w]--; workers[w] == 0 {
				delete(steps, w)
				for _, v := range steps {
					delete(v, w)
				}
				delete(workers, w)
			}
		}
		seconds++
	}
	fmt.Println(seconds)
}

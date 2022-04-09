package main

import (
	"fmt"
	"os"
	"strings"
)

type State struct {
	Pos  string
	Seen map[string]int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	caves := map[string]map[string]struct{}{}
	for _, s := range strings.Fields(string(input)) {
		s := strings.Split(s, "-")

		for a, b := range []int{1, 0} {
			if caves[s[a]] == nil {
				caves[s[a]] = map[string]struct{}{}
			}
			caves[s[a]][s[b]] = struct{}{}
		}
	}

	fmt.Println(run(caves, true))
	fmt.Println(run(caves, false))
}

func run(caves map[string]map[string]struct{}, part1 bool) (count int) {
	queue := []State{{"start", map[string]int{"start": 1}}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Pos == "end" {
			count++
			continue
		}

	out:
		for c := range caves[cur.Pos] {
			if c == "start" {
				continue
			}

			seen := map[string]int{}
			for k, v := range cur.Seen {
				seen[k] = v
				if (part1 || v == 2) && cur.Seen[c] > 0 {
					continue out
				}
			}

			if c == strings.ToLower(c) {
				seen[c]++
			}

			queue = append(queue, State{c, seen})
		}
	}
	return
}

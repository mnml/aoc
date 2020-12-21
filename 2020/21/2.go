package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	count := map[string]int{}
	aller := map[string]map[string]struct{}{}
	for _, s := range split {
		food := regexp.MustCompile(`(.+) \(contains (.+)\)`).FindStringSubmatch(s)

		ingr := map[string]struct{}{}
		for _, s := range strings.Fields(food[1]) {
			count[s]++
			ingr[s] = struct{}{}
		}

		for _, a := range strings.Split(food[2], ", ") {
			if _, ok := aller[a]; !ok {
				aller[a] = map[string]struct{}{}
				for i := range ingr {
					aller[a][i] = struct{}{}
				}
				continue
			}

			for i := range aller[a] {
				if _, ok := ingr[i]; !ok {
					delete(aller[a], i)
				}
			}
		}
	}

	part1 := 0
out:
	for i := range count {
		for _, is := range aller {
			if _, ok := is[i]; ok {
				continue out
			}
		}
		part1 += count[i]
	}
	fmt.Println(part1)

	danger, part2 := map[string]string{}, []string{}
	for len(aller) > 0 {
		for a, is := range aller {
			if len(is) != 1 {
				continue
			}

			for i := range is {
				for _, is := range aller {
					delete(is, i)
				}
				delete(aller, a)
				danger[i] = a
				part2 = append(part2, i)
			}
		}
	}
	sort.Slice(part2, func(i, j int) bool { return danger[part2[i]] < danger[part2[j]] })
	fmt.Println(strings.Join(part2, ","))
}

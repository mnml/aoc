package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	rules := map[string]string{}
	for _, s := range split[2:] {
		split := strings.Split(s, " -> ")
		rules[split[0]] = split[1]
	}

	pairs := map[string]int{}
	for i := 0; i < len(split[0])-1; i++ {
		pairs[split[0][i:i+2]]++
	}

	for i := 0; i < 40; i++ {
		np := map[string]int{}
		for k, v := range pairs {
			np[k[:1]+rules[k]] += v
			np[rules[k]+k[1:]] += v
		}
		pairs = np

		if i == 9 || i == 39 {
			counts := map[byte]int{split[0][len(split[0])-1]: 1}
			for k, v := range pairs {
				counts[k[0]] += v
			}

			vals := []int{}
			for _, v := range counts {
				vals = append(vals, v)
			}
			sort.Sort(sort.IntSlice(vals))
			fmt.Println(vals[len(vals)-1] - vals[0])
		}
	}
}

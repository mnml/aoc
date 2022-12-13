package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	pkts, part1 := []any{}, 0
	for i, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		pkts = append(pkts, a, b)

		if cmp(a, b) <= 0 {
			part1 += i + 1
		}
	}
	fmt.Println(part1)

	pkts = append(pkts, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(pkts, func(i, j int) bool { return cmp(pkts[i], pkts[j]) < 0 })

	part2 := 1
	for i, p := range pkts {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			part2 *= i + 1
		}
	}
	fmt.Println(part2)
}

func cmp(a, b any) int {
	var as, bs []any
	af, bf := false, false

	switch a.(type) {
	case float64:
		as, af = []any{a}, true
	case []any, []float64:
		as = a.([]any)
	}

	switch b.(type) {
	case float64:
		bs, bf = []any{b}, true
	case []any, []float64:
		bs = b.([]any)
	}

	if af && bf {
		return int(as[0].(float64) - bs[0].(float64))
	}
	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}

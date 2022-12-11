package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Monkey struct {
	Items []int
	Op    func(int) int
	Test  func(int) int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	monkeys, lcm := make([]Monkey, len(split)), 1
	for _, s := range split {
		var items, op string
		var i, v, test, t, f int
		fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(s),
			`Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`,
			&i, &items, &op, &v, &test, &t, &f)

		json.Unmarshal([]byte("["+items+"]"), &monkeys[i].Items)
		monkeys[i].Op = map[string]func(int) int{
			"+": func(o int) int { return o + v },
			"*": func(o int) int { return o * v },
			"^": func(o int) int { return o * o },
		}[op]
		monkeys[i].Test = func(w int) int {
			if w%test == 0 {
				return t
			}
			return f
		}
		lcm *= test
	}

	fmt.Println(inspect(monkeys, 20, func(w int) int { return w / 3 }))
	fmt.Println(inspect(monkeys, 10000, func(w int) int { return w % lcm }))
}

func inspect(monkeys []Monkey, rounds int, op func(int) int) int {
	monkeys = append([]Monkey{}, monkeys...)
	inspected := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for i, m := range monkeys {
			for _, w := range m.Items {
				w = op(m.Op(w))
				monkeys[m.Test(w)].Items = append(monkeys[m.Test(w)].Items, w)
				inspected[i]++
			}
			monkeys[i].Items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return inspected[0] * inspected[1]
}

package main

import (
	"fmt"
	"os"
	"strings"
)

type Module struct {
	Type   rune
	Dests  []string
	Memory map[string]bool
}

type Pulse struct {
	Source string
	Value  bool
	Dest   string
}

func main() {
	input, _ := os.ReadFile("input.txt")

	modules := map[string]Module{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		split := strings.Split(s, " -> ")

		name := split[0][1:]
		if split[0] == "broadcaster" {
			name = split[0]
		}

		modules[name] = Module{
			Type:   rune(split[0][0]),
			Dests:  strings.Split(split[1], ", "),
			Memory: map[string]bool{},
		}
	}

	inputs := func(src string) (s []string) {
		for k, m := range modules {
			for _, d := range m.Dests {
				if d == src {
					s = append(s, k)
				}
			}
		}
		return
	}

	counts, cycles := map[bool]int{}, map[string]int{}
	for i := 1; len(cycles) != len(inputs(inputs("rx")[0])); i++ {
		pulses := []Pulse{{"button", false, "broadcaster"}}
		for len(pulses) > 0 {
			p := pulses[0]
			pulses, counts[p.Value] = pulses[1:], counts[p.Value]+1

			m := modules[p.Dest]
			switch m.Type {
			case '%':
				if p.Value {
					continue
				}
				m.Memory[p.Dest] = !m.Memory[p.Dest]
			case '&':
				m.Memory[p.Source], m.Memory[p.Dest] = p.Value, false
				for _, i := range inputs(p.Dest) {
					if !m.Memory[i] {
						m.Memory[p.Dest] = true
					}
				}
			}

			modules[p.Dest] = m
			for _, d := range m.Dests {
				pulses = append(pulses, Pulse{p.Dest, m.Memory[p.Dest], d})
			}

			for _, k := range inputs(inputs("rx")[0]) {
				if _, ok := cycles[k]; !ok && modules[inputs("rx")[0]].Memory[k] {
					cycles[k] = i
				}
			}
		}

		if i == 1000 {
			fmt.Println(counts[false] * counts[true])
		}
	}

	part2 := 1
	for _, v := range cycles {
		part2 *= v
	}
	fmt.Println(part2)
}

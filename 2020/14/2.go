package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	var mask string
	mem1, part1 := map[int]int{}, 0
	mem2, part2 := map[int]int{}, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if _, err := fmt.Sscanf(s, "mask = %s", &mask); err == nil {
			continue
		}
		var addr, value int
		fmt.Sscanf(s, "mem[%d] = %d", &addr, &value)

		for i, x := 0, strings.Count(mask, "X"); i < 1<<x; i++ {
			mask := strings.NewReplacer("X", "x", "0", "X").Replace(mask)
			for _, r := range fmt.Sprintf("%0*b", x, i) {
				mask = strings.Replace(mask, "x", string(r), 1)
			}

			addr := apply(mask, addr)
			part2, mem2[addr] = part2+value-mem2[addr], value
		}

		value = apply(mask, value)
		part1, mem1[addr] = part1+value-mem1[addr], value
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func apply(mask string, value int) int {
	and, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 0)
	or, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 0)
	return value&int(and) | int(or)
}

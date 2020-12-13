package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	time, _ := strconv.Atoi(lines[0])

	part1, part2, step := math.MaxInt64, 0, 1
	for i, s := range strings.Split(lines[1], ",") {
		bus, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		if bus-time%bus < part1-time%part1 {
			part1 = bus
		}

		for (part2+i)%bus != 0 {
			part2 += step
		}
		step *= bus
	}
	fmt.Println(part1 * (part1 - time%part1))
	fmt.Println(part2)
}

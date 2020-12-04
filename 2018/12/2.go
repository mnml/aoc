package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	var state string
	fmt.Fscanf(file, "initial state: %s\n\n", &state)

	notes := map[string]rune{}
	for {
		var k string
		var v rune
		n, _ := fmt.Fscanf(file, "%s => %c\n", &k, &v)
		if n < 2 {
			break
		}
		notes[k] = v
	}

	seen := map[string]int{}
	for g := 1; ; g++ {
		state = "....." + state + "....."

		new := []rune(state)
		for i := 2; i < len(state)-2; i++ {
			new[i] = notes[state[i-2:i+3]]
		}
		state = string(new)

		sum := 0
		for i, r := range state {
			if r == '#' {
				sum += i - (g * 5)
			}
		}

		if g == 20 {
			fmt.Println(sum)
		}
		if prev, ok := seen[strings.Trim(state, ".")]; ok {
			fmt.Println(sum + (sum-prev)*(50000000000-g))
			break
		}
		seen[strings.Trim(state, ".")] = sum
	}
}

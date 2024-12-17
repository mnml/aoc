package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	m := regexp.MustCompile(`[\d,]+`).FindAllString(string(input), -1)

	a, _ := strconv.Atoi(m[0])
	b, _ := strconv.Atoi(m[1])
	c, _ := strconv.Atoi(m[2])
	var pgm []int
	json.Unmarshal([]byte("["+m[3]+"]"), &pgm)

	out := fmt.Sprint(run(a, b, c, pgm))
	fmt.Println(strings.Trim(strings.Replace(out, " ", ",", -1), "[]"))

	a = 0
	for n := len(pgm) - 1; n >= 0; n-- {
		a <<= 3
		for !slices.Equal(run(a, b, c, pgm), pgm[n:]) {
			a++
		}
	}
	fmt.Println(a)
}

func run(a, b, c int, pgm []int) (out []int) {
	for ip := 0; ip < len(pgm); ip += 2 {
		op, literal := pgm[ip], pgm[ip+1]

		combo := literal
		switch combo {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}

		switch op {
		case 0:
			a >>= combo
		case 1:
			b ^= literal
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				ip = literal - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
	}
	return out
}

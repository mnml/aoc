package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	calc := func(time, dist []string) int {
		r := 1
		for i := range time {
			t, _ := strconv.ParseFloat(time[i], 64)
			d, _ := strconv.ParseFloat(dist[i], 64)
			b := math.Sqrt(math.Pow(t, 2) - 4*d)
			r *= int(math.Ceil((t+b)/2) - math.Floor((t-b)/2) - 1)
		}
		return r
	}

	fmt.Println(calc(strings.Fields(split[0])[1:], strings.Fields(split[1])[1:]))
	fmt.Println(calc(
		[]string{strings.Join(strings.Fields(split[0])[1:], "")},
		[]string{strings.Join(strings.Fields(split[1])[1:], "")},
	))
}

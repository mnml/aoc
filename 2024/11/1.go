package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	stones := map[int]int{}
	for _, s := range strings.Fields(string(input)) {
		n, _ := strconv.Atoi(s)
		stones[n]++
	}
	fmt.Println(run(stones, 25))
	fmt.Println(run(stones, 75))
}

func run(stones map[int]int, blinks int) (r int) {
	for range blinks {
		next := map[int]int{}
		for k, v := range stones {
			if k == 0 {
				next[1] += v
			} else if s := strconv.Itoa(k); len(s)%2 == 0 {
				n1, _ := strconv.Atoi(s[:len(s)/2])
				n2, _ := strconv.Atoi(s[len(s)/2:])
				next[n1] += v
				next[n2] += v
			} else {
				next[k*2024] += v
			}
		}
		stones = next
	}
	for _, v := range stones {
		r += v
	}
	return r
}

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	xmas := make([]int, len(split))
	for i, s := range split {
		xmas[i], _ = strconv.Atoi(s)
	}

	invalid := 0
out:
	for i := 25; i < len(xmas); i++ {
		for j := i - 25; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if xmas[j]+xmas[k] == xmas[i] {
					continue out
				}
			}
		}
		invalid = xmas[i]
		break
	}
	fmt.Println(invalid)

	for i := 0; i < len(xmas); i++ {
		for j := i + 1; j < len(xmas); j++ {
			sum := 0
			for _, v := range xmas[i : j+1] {
				sum += v
			}
			if sum == invalid {
				sort.Ints(xmas[i : j+1])
				fmt.Println(xmas[i] + xmas[j])
				return
			}
		}
	}
}

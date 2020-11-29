package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	str := strings.TrimSpace(string(input))
	w, h := 25, 6
	count := make([]map[rune]int, len(str)/(w*h))
	min := 0

	for l := range count {
		count[l] = map[rune]int{}

		for _, r := range str[l*w*h : (l+1)*w*h] {
			count[l][r]++
		}

		if count[l]['0'] < count[min]['0'] {
			min = l
		}
	}

	fmt.Println(count[min]['1'] * count[min]['2'])
}

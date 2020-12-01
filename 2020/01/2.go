package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	entries := []int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		i, _ := strconv.Atoi(s)
		entries = append(entries, i)
	}

	for i, x := range entries {
		for j, y := range entries[i+1:] {
			if x+y == 2020 {
				fmt.Println(x * y)
			}
			for _, z := range entries[j+1:] {
				if x+y+z == 2020 {
					fmt.Println(x * y * z)
				}
			}
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	total := 0
	reached := map[int]struct{}{}
	for {
		for _, s := range strings.Fields(string(input)) {
			i, _ := strconv.Atoi(s)
			total += i
			if _, ok := reached[total]; ok {
				fmt.Println(total)
				return
			}
			reached[total] = struct{}{}
		}
	}
}

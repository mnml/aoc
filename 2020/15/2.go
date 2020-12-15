package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	spoken, last := map[int]int{}, 0
	for i, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		last, _ = strconv.Atoi(s)
		spoken[last] = i + 1
	}

	for i := len(spoken); i < 30000000; i++ {
		if v, ok := spoken[last]; ok {
			spoken[last], last = i, i-v
		} else {
			spoken[last], last = i, 0
		}

		if i == 2020-1 || i == 30000000-1 {
			fmt.Println(last)
		}
	}
}

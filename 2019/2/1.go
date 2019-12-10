package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(input), ",")
	m := make([]int, len(split))

	for i, s := range split {
		m[i], _ = strconv.Atoi(s)
	}

	m[1], m[2] = 12, 2

	for i := 0; i < len(m); i += 4 {
		switch m[i] {
		case 1:
			m[m[i+3]] = m[m[i+1]] + m[m[i+2]]
		case 2:
			m[m[i+3]] = m[m[i+1]] * m[m[i+2]]
		case 99:
			fmt.Println(m[0])
			i = len(m)
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	m := make([]int, len(split))
	pc := 0

	for i, s := range split {
		m[i], _ = strconv.Atoi(s)
	}

	for {
		ins := fmt.Sprintf("%05d", m[pc])
		op, _ := strconv.Atoi(ins[3:])
		arg := func(i int) int {
			if ins[3-i] == '0' {
				return m[m[pc+i]]
			}
			return m[pc+i]
		}

		switch op {
		case 1:
			m[m[pc+3]] = arg(1) + arg(2)
		case 2:
			m[m[pc+3]] = arg(1) * arg(2)
		case 3:
			fmt.Scan(&m[m[pc+1]])
		case 4:
			fmt.Println(arg(1))
		case 5:
			if arg(1) != 0 {
				pc = arg(2)
				continue
			}
		case 6:
			if arg(1) == 0 {
				pc = arg(2)
				continue
			}
		case 7:
			if arg(1) < arg(2) {
				m[m[pc+3]] = 1
			} else {
				m[m[pc+3]] = 0
			}
		case 8:
			if arg(1) == arg(2) {
				m[m[pc+3]] = 1
			} else {
				m[m[pc+3]] = 0
			}
		case 99:
			return
		}

		pc += []int{1, 4, 4, 2, 2, 3, 3, 4, 4}[op]
	}
}

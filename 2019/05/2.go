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
	mem := make([]int, len(split))
	ip := 0

	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		arg := func(i int) int {
			if ins[3-i] == '0' {
				return mem[mem[ip+i]]
			}
			return mem[ip+i]
		}

		switch op {
		case 1:
			mem[mem[ip+3]] = arg(1) + arg(2)
		case 2:
			mem[mem[ip+3]] = arg(1) * arg(2)
		case 3:
			fmt.Scan(&mem[mem[ip+1]])
		case 4:
			fmt.Println(arg(1))
		case 5:
			if arg(1) != 0 {
				ip = arg(2)
				continue
			}
		case 6:
			if arg(1) == 0 {
				ip = arg(2)
				continue
			}
		case 7:
			if arg(1) < arg(2) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
		case 8:
			if arg(1) == arg(2) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
		case 99:
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4}[op]
	}
}

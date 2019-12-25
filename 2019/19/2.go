package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	mem := map[int]int{}
	for i, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		mem[i], _ = strconv.Atoi(s)
	}

	x, y := 0, 99
	for {
		for run(mem, []int{x, y})[0] != 1 {
			x++
		}
		if run(mem, []int{x + 99, y - 99})[0] == 1 {
			fmt.Println(x*10000 + y - 99)
			return
		}
		y++
	}
}

func run(init map[int]int, in []int) (out []int) {
	ip, rb := 0, 0
	mem := map[int]int{}
	for i, v := range init {
		mem[i] = v
	}

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		par := func(i int) int {
			switch ins[3-i] {
			case '1':
				return ip + i
			case '2':
				return rb + mem[ip+i]
			default:
				return mem[ip+i]
			}
		}

		switch op {
		case 1:
			mem[par(3)] = mem[par(1)] + mem[par(2)]
		case 2:
			mem[par(3)] = mem[par(1)] * mem[par(2)]
		case 3:
			mem[par(1)] = in[0]
			in = in[1:]
		case 4:
			out = append(out, mem[par(1)])
		case 5:
			if mem[par(1)] != 0 {
				ip = mem[par(2)]
				continue
			}
		case 6:
			if mem[par(1)] == 0 {
				ip = mem[par(2)]
				continue
			}
		case 7:
			if mem[par(1)] < mem[par(2)] {
				mem[par(3)] = 1
			} else {
				mem[par(3)] = 0
			}
		case 8:
			if mem[par(1)] == mem[par(2)] {
				mem[par(3)] = 1
			} else {
				mem[par(3)] = 0
			}
		case 9:
			rb += mem[par(1)]
		case 99:
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}

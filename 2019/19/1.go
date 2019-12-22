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

	sum := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			sum += run(mem, []int{x, y})[0]
		}
	}

	fmt.Println(sum)
}

func run(init map[int]int, in []int) (out []int) {
	mem := map[int]int{}
	for i, v := range init {
		mem[i] = v
	}

	ip := 0
	rb := 0

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		par := func(i int) (addr int) {
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

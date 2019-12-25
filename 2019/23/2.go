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

	in, out := make([]chan int, 50), make([]chan int, 50)
	for i := 0; i < 50; i++ {
		in[i], out[i] = make(chan int), make(chan int)
		go run(mem, in[i], out[i])
		in[i] <- i
		in[i] <- -1
	}

	idle := 0
	var old, nat [2]int

	for i := 0; ; i = (i + 1) % 50 {
		select {
		case addr := <-out[i]:
			if addr == 255 {
				new := [2]int{<-out[i], <-out[i]}
				if nat == [2]int{} {
					fmt.Println(new[1])
				}
				nat = new
			} else {
				in[addr] <- <-out[i]
				in[addr] <- <-out[i]
			}
			idle = 0
		case in[i] <- -1:
			idle++
		}

		if idle >= 50 {
			if nat[1] == old[1] {
				fmt.Println(nat[1])
				return
			}
			in[0] <- nat[0]
			in[0] <- nat[1]
			old = nat
			idle = 0
		}
	}
}

func run(init map[int]int, in <-chan int, out chan<- int) {
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
			mem[par(1)] = <-in
		case 4:
			out <- mem[par(1)]
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

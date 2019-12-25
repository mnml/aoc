package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")

	mem := map[int]int{}
	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	in, out := make(chan int, 1), make(chan int)
	go run(mem, in, out)
	hull := map[image.Point]int{}
	pos, dir := image.Point{}, 0

	in <- 1
	for hull[pos] = range out {
		dir = (dir + 2*<-out + 1) % 4
		pos = pos.Add([]image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}[dir])
		in <- hull[pos]
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			fmt.Print([]string{"  ", "██"}[hull[image.Point{x, y}]])
		}
		fmt.Println()
	}
}

func run(mem map[int]int, in <-chan int, out chan<- int) {
	ip, rb := 0, 0

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
			close(out)
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}

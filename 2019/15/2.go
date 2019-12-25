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

	in, out := make(chan int), make(chan int)
	go run(mem, in, out)
	delta := map[int]image.Point{1: {0, -1}, 2: {0, 1}, 3: {-1, 0}, 4: {1, 0}}
	maze := map[image.Point]int{}
	pos := image.Point{0, 0}
	var back []int
	var oxy image.Point

explore:
	for {
		for dir, dp := range delta {
			next := pos.Add(dp)
			if _, ok := maze[next]; !ok {
				in <- dir
				maze[next] = <-out
				if maze[next] > 0 {
					pos = next
					back = append(back, dir-1^1+1)
				}
				if maze[next] == 2 {
					oxy = next
					fmt.Println(len(back))
				}
				continue explore
			}
		}
		if len(back) < 1 {
			break
		}
		in <- back[len(back)-1]
		<-out
		pos = pos.Add(delta[back[len(back)-1]])
		back = back[:len(back)-1]
	}

	disc := map[image.Point]int{oxy: 0}
	queue := []image.Point{oxy}
	var point image.Point

	for len(queue) > 0 {
		point = queue[0]
		queue = queue[1:]
		for _, d := range delta {
			if _, ok := disc[point.Add(d)]; !ok && maze[point.Add(d)] > 0 {
				disc[point.Add(d)] = disc[point] + 1
				queue = append(queue, point.Add(d))
			}
		}
	}

	fmt.Println(disc[point])
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
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}

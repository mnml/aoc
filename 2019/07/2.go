package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	seq := []int{5, 6, 7, 8, 9}

	mem := make([]int, len(split))
	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	max := 0
	for _, p := range perms(seq) {
		chans := make([]chan int, len(seq))
		for i := range seq {
			chans[i] = make(chan int, 1)
		}

		var wg sync.WaitGroup
		for i, v := range p {
			wg.Add(1)
			go run(append([]int(nil), mem...), chans[i], chans[(i+1)%len(seq)], &wg)
			chans[i] <- v
		}

		chans[0] <- 0
		wg.Wait()
		max = int(math.Max(float64(<-chans[0]), float64(max)))
	}

	fmt.Println(max)
}

func run(mem []int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	ip := 0

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		par := func(i int) int {
			if ins[3-i] == '0' {
				return mem[mem[ip+i]]
			}
			return mem[ip+i]
		}

		switch op {
		case 1:
			mem[mem[ip+3]] = par(1) + par(2)
		case 2:
			mem[mem[ip+3]] = par(1) * par(2)
		case 3:
			mem[mem[ip+1]] = <-in
		case 4:
			out <- par(1)
		case 5:
			if par(1) != 0 {
				ip = par(2)
				continue
			}
		case 6:
			if par(1) == 0 {
				ip = par(2)
				continue
			}
		case 7:
			if par(1) < par(2) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
		case 8:
			if par(1) == par(2) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
		case 99:
			wg.Done()
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4}[op]
	}
}

func perms(ints []int) [][]int {
	out := [][]int{}

	if len(ints) == 1 {
		return [][]int{ints}
	}

	for i := range ints {
		c := append([]int(nil), ints...)
		for _, p := range perms(append(c[:i], c[i+1:]...)) {
			out = append(out, append([]int{ints[i]}, p...))
		}
	}

	return out
}

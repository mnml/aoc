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
	pc := 0

	for {
		ins := fmt.Sprintf("%05d", mem[pc])
		op, _ := strconv.Atoi(ins[3:])
		arg := func(i int) int {
			if ins[3-i] == '0' {
				return mem[mem[pc+i]]
			}
			return mem[pc+i]
		}

		switch op {
		case 1:
			mem[mem[pc+3]] = arg(1) + arg(2)
		case 2:
			mem[mem[pc+3]] = arg(1) * arg(2)
		case 3:
			mem[mem[pc+1]] = <-in
		case 4:
			out <- arg(1)
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
				mem[mem[pc+3]] = 1
			} else {
				mem[mem[pc+3]] = 0
			}
		case 8:
			if arg(1) == arg(2) {
				mem[mem[pc+3]] = 1
			} else {
				mem[mem[pc+3]] = 0
			}
		case 99:
			wg.Done()
			return
		}

		pc += []int{1, 4, 4, 2, 2, 3, 3, 4, 4}[op]
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

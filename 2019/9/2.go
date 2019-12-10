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
	mem := map[int]int64{}
	pc := 0
	base := 0

	for i, s := range split {
		mem[i], _ = strconv.ParseInt(s, 10, 0)
	}

	for {
		ins := fmt.Sprintf("%05d", mem[pc])
		op, _ := strconv.Atoi(ins[3:])
		arg := func(i int) (addr int) {
			switch ins[3-i] {
			case '1':
				return pc + i
			case '2':
				return base + int(mem[pc+i])
			default:
				return int(mem[pc+i])
			}
		}

		switch op {
		case 1:
			mem[arg(3)] = mem[arg(1)] + mem[arg(2)]
		case 2:
			mem[arg(3)] = mem[arg(1)] * mem[arg(2)]
		case 3:
			var i int64
			fmt.Scan(&i)
			mem[arg(1)] = i
		case 4:
			fmt.Println(mem[arg(1)])
		case 5:
			if mem[arg(1)] != 0 {
				pc = int(mem[arg(2)])
				continue
			}
		case 6:
			if mem[arg(1)] == 0 {
				pc = int(mem[arg(2)])
				continue
			}
		case 7:
			if mem[arg(1)] < mem[arg(2)] {
				mem[arg(3)] = 1
			} else {
				mem[arg(3)] = 0
			}
		case 8:
			if mem[arg(1)] == mem[arg(2)] {
				mem[arg(3)] = 1
			} else {
				mem[arg(3)] = 0
			}
		case 9:
			base += int(mem[arg(1)])
		case 99:
			return
		}

		pc += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}

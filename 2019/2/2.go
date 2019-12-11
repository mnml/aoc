package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(input), ",")

	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			mem := make([]int, len(split))

			for i, s := range split {
				mem[i], _ = strconv.Atoi(s)
			}

			mem[1], mem[2] = n, v

			for i := 0; i < len(mem); i += 4 {
				switch mem[i] {
				case 1:
					mem[mem[i+3]] = mem[mem[i+1]] + mem[mem[i+2]]
				case 2:
					mem[mem[i+3]] = mem[mem[i+1]] * mem[mem[i+2]]
				case 99:
					if mem[0] == 19690720 {
						fmt.Println(100*n + v)
					}
					i = len(mem)
				}
			}
		}
	}
}

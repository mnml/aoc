package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	ins := strings.Split(strings.TrimSpace(string(input)), "\n")

	acc, _ := run(ins)
	fmt.Println(acc)

	for i, s := range ins {
		tmp := make([]string, len(ins))
		copy(tmp, ins)
		tmp[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(s)

		if acc, err := run(tmp); err == nil {
			fmt.Println(acc)
			break
		}
	}
}

func run(ins []string) (int, error) {
	pc, acc := 0, 0
	seen := map[int]struct{}{}

	for pc < len(ins) {
		if _, ok := seen[pc]; ok {
			return acc, errors.New("infinite loop")
		}
		seen[pc] = struct{}{}

		var op string
		var arg int
		fmt.Sscanf(ins[pc], "%s %d", &op, &arg)

		switch op {
		case "acc":
			acc += arg
		case "jmp":
			pc += arg - 1
		}
		pc++
	}

	return acc, nil
}

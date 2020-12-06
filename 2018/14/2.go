package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 909441

func main() {
	instr := strconv.Itoa(input)
	elf1, elf2 := 0, 1

	scores := []byte{'3', '7'}
	for len(scores) < len(instr) || string(scores[len(scores)-len(instr):]) != instr {
		scores = append(scores, []byte(strconv.Itoa(int(scores[elf1]-'0'+scores[elf2]-'0')))...)
		elf1 = (elf1 + 1 + int(scores[elf1]-'0')) % len(scores)
		elf2 = (elf2 + 1 + int(scores[elf2]-'0')) % len(scores)
	}
	fmt.Println(string(scores[input : input+10]))
	fmt.Println(strings.Index(string(scores), instr))
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fields := strings.Fields(string(input))

	gamma, epsilon := "", ""
	oxygen, co2 := fields, fields
	for i := range fields[0] {
		most, least := split(fields, i)
		gamma += string(most[0][i])
		epsilon += string(least[0][i])

		if len(oxygen) > 1 {
			oxygen, _ = split(oxygen, i)
		}
		if len(co2) > 1 {
			_, co2 = split(co2, i)
		}
	}
	fmt.Println(bs2i(gamma) * bs2i(epsilon))
	fmt.Println(bs2i(oxygen[0]) * bs2i(co2[0]))
}

func split(report []string, index int) (most, least []string) {
	bit := map[byte][]string{}
	for _, s := range report {
		bit[s[index]] = append(bit[s[index]], s)
	}

	if len(bit['1']) >= len(bit['0']) {
		return bit['1'], bit['0']
	}
	return bit['0'], bit['1']
}

func bs2i(bs string) int {
	i, _ := strconv.ParseUint(bs, 2, len(bs))
	return int(i)
}

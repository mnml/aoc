package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		part1 += strings.Index("B XC YA ZA XB YC ZC XA YB Z", s)/3 + 1
		part2 += strings.Index("B XC XA XA YB YC YC ZA ZB Z", s)/3 + 1
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

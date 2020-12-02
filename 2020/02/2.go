package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var min, max int
		var char byte
		var pass string
		fmt.Sscanf(s, "%v-%v %c: %v", &min, &max, &char, &pass)
		count := strings.Count(pass, string(char))
		if count >= min && count <= max {
			part1++
		}
		if (pass[min-1] == char) != (pass[max-1] == char) {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

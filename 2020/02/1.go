package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	valid := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var min, max int
		var char, pass string
		fmt.Sscanf(s, "%d-%d %1s: %s", &min, &max, &char, &pass)
		count := strings.Count(pass, char)
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Println(valid)
}

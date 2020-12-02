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
		var char byte
		var pass string
		fmt.Sscanf(s, "%d-%d %c: %s", &min, &max, &char, &pass)
		if (pass[min-1] == char) != (pass[max-1] == char) {
			valid++
		}
	}
	fmt.Println(valid)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	total := 0
	for _, s := range strings.Fields(string(input)) {
		i, _ := strconv.Atoi(s)
		total += i
	}
	fmt.Println(total)
}

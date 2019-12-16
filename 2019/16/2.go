package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	signal := strings.Repeat(strings.TrimSpace(string(input)), 10000)
	offset, _ := strconv.Atoi(signal[:7])
	output := []int{}

	for _, r := range signal[offset:] {
		output = append(output, int(r-'0'))
	}

	for p := 0; p < 100; p++ {
		sum := 0
		for i := len(output) - 1; i >= 0; i-- {
			sum += output[i]
			output[i] = sum % 10
		}
	}

	for _, i := range output[:8] {
		fmt.Print(i)
	}
	fmt.Println()
}

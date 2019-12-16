package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	signal := strings.TrimSpace(string(input))

	for p := 0; p < 100; p++ {
		output := ""
		for i := range signal {
			sum := 0
			for j, r := range signal {
				sum += int(r-'0') * []int{0, 1, 0, -1}[(j+1)/(i+1)%4]
			}
			if sum < 0 {
				sum = -sum
			}
			output += string('0' + sum%10)
		}
		signal = output
	}

	fmt.Println(signal[:8])
}

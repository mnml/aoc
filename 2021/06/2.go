package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fish := make([]int, 9)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		d, _ := strconv.Atoi(s)
		fish[d]++
	}

	for i := 0; i < 256; i++ {
		if i == 80 {
			fmt.Println(sum(fish))
		}
		fish = append(fish[1:7], fish[7]+fish[0], fish[8], fish[0])
	}
	fmt.Println(sum(fish))
}

func sum(xs []int) (sum int) {
	for _, x := range xs {
		sum += x
	}
	return
}

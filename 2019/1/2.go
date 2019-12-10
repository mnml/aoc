package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fuel(mass int) int {
	return mass/3 - 2
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())

		for f := fuel(mass); f > 0; f = fuel(f) {
			total += f
		}
	}

	fmt.Println(total)
}

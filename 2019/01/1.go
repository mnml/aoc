package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += mass/3 - 2
	}

	fmt.Println(total)
}

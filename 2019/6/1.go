package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	orbits := map[string]string{}
	total := 0

	for scanner.Scan() {
		obj := strings.Split(scanner.Text(), ")")
		orbits[obj[1]] = obj[0]
	}

	for o := range orbits {
		for o, ok := orbits[o]; ok; o, ok = orbits[o] {
			total++
		}
	}

	fmt.Println(total)
}

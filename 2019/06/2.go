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

	for scanner.Scan() {
		objs := strings.Split(scanner.Text(), ")")
		orbits[objs[1]] = objs[0]
	}

	you := map[string]int{}
	for o, ok := orbits["YOU"]; ok; o, ok = orbits[o] {
		you[o] = len(you)
	}

	san := 0
	for o, ok := orbits["SAN"]; ok; o, ok = orbits[o] {
		if _, ok := you[o]; ok {
			fmt.Println(san + you[o])
			break
		}
		san++
	}
}

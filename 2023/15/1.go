package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`(\w+)([-=])(\d*)`)

	part1, part2 := 0, 0
	boxes, focal := [256][]string{}, map[string]int{}
	for _, m := range re.FindAllStringSubmatch(string(input), -1) {
		h := hash(m[1])
		i := slices.Index(boxes[h], m[1])

		if m[2] == "-" && i != -1 {
			boxes[h] = slices.Delete(boxes[h], i, i+1)
		} else if m[2] == "=" {
			focal[m[1]] = int(m[3][0] - '0')
			if i == -1 {
				boxes[h] = append(boxes[h], m[1])
			}
		}

		part1 += hash(m[0])
	}

	for i, b := range boxes {
		for j, l := range b {
			part2 += (i + 1) * (j + 1) * focal[l]
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func hash(s string) (h int) {
	for _, r := range s {
		h = (h + int(r)) * 17 % 256
	}
	return
}

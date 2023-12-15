package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
)

type Lens struct {
	Label string
	Focal int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`(\w+)([-=])(\d*)`)

	part1, part2 := 0, 0
	boxes := [256][]Lens{}
	for _, m := range re.FindAllStringSubmatch(string(input), -1) {
		box := &boxes[hash(m[1])]
		i := slices.IndexFunc(*box, func(l Lens) bool { return l.Label == m[1] })

		if m[2] == "-" && i != -1 {
			*box = slices.Delete(*box, i, i+1)
		} else if m[2] == "=" && i != -1 {
			(*box)[i] = Lens{m[1], int(m[3][0] - '0')}
		} else if m[2] == "=" {
			*box = append(*box, Lens{m[1], int(m[3][0] - '0')})
		}

		part1 += hash(m[0])
	}

	for i, b := range boxes {
		for j, l := range b {
			part2 += (i + 1) * (j + 1) * l.Focal
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

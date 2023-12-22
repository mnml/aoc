package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Brick struct {
	x1, y1, z1 int
	x2, y2, z2 int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	bricks := []Brick{}
	for _, s := range strings.Fields(string(input)) {
		var b Brick
		fmt.Sscanf(s, "%d,%d,%d~%d,%d,%d", &b.x1, &b.y1, &b.z1, &b.x2, &b.y2, &b.z2)
		bricks = append(bricks, b)
	}
	slices.SortFunc(bricks, func(a, b Brick) int { return cmp.Compare(a.z1, b.z1) })

	for n, ok := 0, true; ok; ok = n != 0 {
		bricks, n = drop(bricks)
	}

	part1, part2 := 0, 0
	for i := range bricks {
		_, n := drop(slices.Delete(slices.Clone(bricks), i, i+1))
		if n == 0 {
			part1++
		}
		part2 += n
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func drop(bricks []Brick) (bs []Brick, n int) {
	bs = slices.Clone(bricks)
loop:
	for i, a := range bs {
		if a.z1 == 1 || a.z2 == 1 {
			continue
		}
		a.z1, a.z2 = a.z1-1, a.z2-1

		for _, b := range bs[:i] {
			if a.x1 <= b.x2 && a.x2 >= b.x1 &&
				a.y1 <= b.y2 && a.y2 >= b.y1 &&
				a.z1 <= b.z2 && a.z2 >= b.z1 {
				continue loop
			}
		}
		bs[i] = a
		n++
	}
	return
}

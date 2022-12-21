package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println(mix(string(input), 1, 1))
	fmt.Println(mix(string(input), 811589153, 10))
}

func mix(input string, key, times int) int {
	split := strings.Fields(input)

	r, idx, z := ring.New(len(split)), map[int]*ring.Ring{}, (*ring.Ring)(nil)
	for i, s := range split {
		v, _ := strconv.Atoi(s)
		if v == 0 {
			z = r
		}
		r.Value, idx[i], r = v*key, r, r.Next()
	}

	for i := 0; i < times; i++ {
		for i := 0; i < len(idx); i++ {
			r = idx[i].Prev()
			c := r.Unlink(1)
			r.Move(c.Value.(int) % (len(idx) - 1)).Link(c)
		}
	}

	return z.Move(1000).Value.(int) + z.Move(2000).Value.(int) + z.Move(3000).Value.(int)
}

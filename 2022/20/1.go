package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	vs := [][2]int{}
	for i, s := range strings.Fields(string(input)) {
		v, _ := strconv.Atoi(s)
		vs = append(vs, [2]int{i, v})
	}

	fmt.Println(mix(vs, 1, 1))
	fmt.Println(mix(vs, 811589153, 10))
}

func mix(vs [][2]int, key, times int) int {
	vs = append([][2]int(nil), vs...)
	for i := range vs {
		vs[i][1] *= key
	}

	for t := 0; t < times; t++ {
		for i := 0; i < len(vs); i++ {
			i := index(vs, func(v [2]int) bool { return v[0] == i })
			v := vs[i]
			vs = append(vs[:i], vs[i+1:]...)
			i = mod(i+v[1], len(vs))
			vs = append(vs[:i], append([][2]int{v}, vs[i:]...)...)
		}
	}

	i := index(vs, func(v [2]int) bool { return v[1] == 0 })
	return vs[mod(i+1000, len(vs))][1] + vs[mod(i+2000, len(vs))][1] + vs[mod(i+3000, len(vs))][1]
}

func index[T any](s []T, f func(T) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

func mod(a, n int) int {
	return (a%n + n) % n
}

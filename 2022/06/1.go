package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(find(string(input), 4))
	fmt.Println(find(string(input), 14))
}

func find(s string, l int) int {
	for i := l; i <= len(s); i++ {
		m := map[rune]struct{}{}
		for _, r := range s[i-l : i] {
			m[r] = struct{}{}
		}
		if len(m) >= l {
			return i
		}
	}
	return -1
}

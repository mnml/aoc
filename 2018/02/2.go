package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	seen := map[string]struct{}{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for i := range s {
			t := s[:i] + " " + s[i+1:]
			if _, ok := seen[t]; ok {
				fmt.Println(strings.ReplaceAll(t, " ", ""))
				return
			}
			seen[t] = struct{}{}
		}
	}
}

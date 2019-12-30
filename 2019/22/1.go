package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	deck := make([]int, 10007)
	for i := 0; i < len(deck); i++ {
		deck[i] = i
	}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		switch s := strings.Fields(s); s[1] {
		case "into":
			for i := len(deck)/2 - 1; i >= 0; i-- {
				deck[i], deck[len(deck)-1-i] = deck[len(deck)-1-i], deck[i]
			}
		case "with":
			inc, _ := strconv.Atoi(s[3])
			table := make([]int, len(deck))
			for i := 0; i < len(deck); i++ {
				table[i*inc%len(deck)] = deck[i]
			}
			deck = table
		default:
			cut, _ := strconv.Atoi(s[1])
			if cut > 0 {
				deck = append(deck[cut:], deck[:cut]...)
			} else {
				deck = append(deck[len(deck)+cut:], deck[:len(deck)+cut]...)
			}
		}
	}

	for i, v := range deck {
		if v == 2019 {
			fmt.Println(i)
			return
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Reac struct {
	Out int
	In  map[string]int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	reacs := map[string]Reac{}

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		reac := strings.Split(s, " => ")
		out := strings.Split(reac[1], " ")
		amount, _ := strconv.Atoi(out[0])
		reacs[out[1]] = Reac{amount, map[string]int{}}

		for _, s := range strings.Split(reac[0], ", ") {
			in := strings.Split(s, " ")
			amount, _ := strconv.Atoi(in[0])
			reacs[out[1]].In[in[1]] = amount
		}
	}

	fmt.Println(ore(map[string]int{"FUEL": 1}, reacs))

	fmt.Println(sort.Search(1000000000000, func(n int) bool {
		return ore(map[string]int{"FUEL": n}, reacs) > 1000000000000
	}) - 1)
}

func ore(want map[string]int, reacs map[string]Reac) int {
loop:
	for {
		for w := range want {
			if w != "ORE" && want[w] > 0 {
				amount := (want[w]-1)/reacs[w].Out + 1
				want[w] -= reacs[w].Out * amount

				for r := range reacs[w].In {
					want[r] += reacs[w].In[r] * amount
				}
				continue loop
			}
		}
		return want["ORE"]
	}
}

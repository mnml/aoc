package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	decks := make([][]int, len(split))
	for i, s := range split {
		for _, s := range strings.Split(s, "\n")[1:] {
			c, _ := strconv.Atoi(s)
			decks[i] = append(decks[i], c)
		}
	}

	_, score := run([][]int{append([]int{}, decks[0]...), append([]int{}, decks[1]...)}, false)
	fmt.Println(score)
	_, score = run(decks, true)
	fmt.Println(score)
}

func run(ds [][]int, rec bool) (win int, score int) {
	seen := map[string]struct{}{}

	for len(ds[0]) > 0 && len(ds[1]) > 0 {
		win = 0
		if _, ok := seen[fmt.Sprint(ds)]; rec && ok {
			break
		}
		seen[fmt.Sprint(ds)] = struct{}{}

		if rec && len(ds[0]) > ds[0][0] && len(ds[1]) > ds[1][0] {
			win, _ = run([][]int{append([]int{}, ds[0][1:ds[0][0]+1]...), append([]int{}, ds[1][1:ds[1][0]+1]...)}, rec)
		} else if ds[0][0] < ds[1][0] {
			win = 1
		}

		ds[win] = append(ds[win], ds[win][0], ds[-win+1][0])
		ds[0], ds[1] = ds[0][1:], ds[1][1:]
	}

	for i, c := range ds[win] {
		score += c * (len(ds[win]) - i)
	}
	return
}

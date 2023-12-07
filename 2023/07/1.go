package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	hands := []Hand{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		h := Hand{}
		fmt.Sscanf(s, "%s %d", &h.Cards, &h.Bid)
		hands = append(hands, h)
	}

	winnings := func(jokers bool) (w int) {
		slices.SortFunc(hands, func(a, b Hand) int {
			return cmp(a.Cards, b.Cards, jokers)
		})
		for i, h := range hands {
			w += (i + 1) * h.Bid
		}
		return
	}

	fmt.Println(winnings(false))
	fmt.Println(winnings(true))
}

func cmp(a, b string, jokers bool) int {
	j, r := "J", "TAJBQCKDAE"
	if jokers {
		j, r = "23456789TQKA", "TAJ0QCKDAE"
	}

	typ := func(cards string) string {
		k := 0
		for _, j := range strings.Split(j, "") {
			n, t := strings.ReplaceAll(cards, "J", j), 0
			for _, s := range n {
				t += strings.Count(n, string(s))
			}
			k = slices.Max([]int{k, t})
		}
		return map[int]string{5: "0", 7: "1", 9: "2", 11: "3", 13: "4", 17: "5", 25: "6"}[k]
	}

	return strings.Compare(
		typ(a)+strings.NewReplacer(strings.Split(r, "")...).Replace(a),
		typ(b)+strings.NewReplacer(strings.Split(r, "")...).Replace(b),
	)
}

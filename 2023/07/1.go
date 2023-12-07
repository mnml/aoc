package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
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
			return rank(a.Cards, jokers) - rank(b.Cards, jokers)
		})
		for i, h := range hands {
			w += (i + 1) * h.Bid
		}
		return
	}

	fmt.Println(winnings(false))
	fmt.Println(winnings(true))
}

func rank(cards string, jokers bool) int {
	j, r := "J", "2031425364758697T8J9QAKBAC"
	if jokers {
		j, r = "23456789TQKA", "J02132435465768798T9QAKBAC"
	}

	typ := 0
	for _, j := range strings.Split(j, "") {
		n, t := strings.ReplaceAll(cards, "J", j), 0
		for _, s := range n {
			t += strings.Count(n, string(s))
		}
		typ = slices.Max([]int{typ, t})
	}

	tie, _ := strconv.ParseInt(strings.NewReplacer(
		strings.Split(r, "")...,
	).Replace(cards), 13, strconv.IntSize)

	return int(math.Pow(13, float64(len(cards))))*typ + int(tie)
}

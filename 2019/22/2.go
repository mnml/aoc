package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	cards, shuff := big.NewInt(119315717514047), big.NewInt(101741582076661)
	a, b := big.NewInt(1), big.NewInt(0)

	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		switch s := strings.Fields(s); s[1] {
		case "into":
			a.Neg(a)
			b.Add(b, a)
		case "with":
			inc, _ := new(big.Int).SetString(s[3], 10)
			a.Mul(a, new(big.Int).ModInverse(inc, cards))
		default:
			cut, _ := new(big.Int).SetString(s[1], 10)
			b.Add(b, new(big.Int).Mul(a, cut))
		}
	}

	b.Mul(b, new(big.Int).Sub(big.NewInt(1), new(big.Int).Exp(a, shuff, cards)))
	b.Mul(b, new(big.Int).ModInverse(new(big.Int).Sub(big.NewInt(1), a), cards))
	a.Exp(a, shuff, cards)

	fmt.Println(new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(a, big.NewInt(2020)), b), cards))
}

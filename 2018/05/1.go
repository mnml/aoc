package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	polymer := strings.TrimSpace(string(input))
	for {
		old := polymer
		for r := 'a'; r <= 'z'; r++ {
			polymer = strings.ReplaceAll(polymer, string([]rune{r, unicode.ToUpper(r)}), "")
			polymer = strings.ReplaceAll(polymer, string([]rune{unicode.ToUpper(r), r}), "")
		}
		if polymer == old {
			fmt.Println(len(polymer))
			return
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	min := int(^uint(0) >> 1)
	for r := 'a'; r <= 'z'; r++ {
		polymer := strings.TrimSpace(string(input))
		polymer = strings.ReplaceAll(polymer, string(r), "")
		polymer = strings.ReplaceAll(polymer, string(unicode.ToUpper(r)), "")

		for {
			old := polymer
			for r := 'a'; r <= 'z'; r++ {
				polymer = strings.ReplaceAll(polymer, string([]rune{r, unicode.ToUpper(r)}), "")
				polymer = strings.ReplaceAll(polymer, string([]rune{unicode.ToUpper(r), r}), "")
			}
			if polymer == old {
				if len(polymer) < min {
					min = len(polymer)
				}
				break
			}
		}
	}
	fmt.Println(min)
}

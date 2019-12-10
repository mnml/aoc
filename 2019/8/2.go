package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	str := strings.TrimSpace(string(input))
	w, h := 25, 6

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			for l := 0; l < len(str)/(w*h); l++ {
				if r := str[l*w*h+y*w+x]; r != '2' {
					fmt.Print(map[byte]string{'0': "  ", '1': "██"}[r])
					break
				}
			}
		}
		fmt.Println()
	}
}

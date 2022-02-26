package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	re := regexp.MustCompile(`\d+`)

	boards := [][]int{}
	for i, s := range split {
		boards = append(boards, []int{})
		for _, s := range re.FindAllString(s, -1) {
			d, _ := strconv.Atoi(s)
			boards[i] = append(boards[i], d)
		}
	}

	scores := []int{}
	for _, d := range boards[0] {
		for i := len(boards) - 1; i >= 1; i-- {
			b := boards[i]

			sum := 0
			for i, n := range b {
				if n == d {
					b[i] = ^n
				} else if n > 0 {
					sum += n
				}
			}

			for j := 0; j < 5; j++ {
				if b[j*5]&b[j*5+1]&b[j*5+2]&b[j*5+3]&b[j*5+4]|
					b[j]&b[j+5]&b[j+10]&b[j+15]&b[j+20] < 0 {
					scores = append(scores, sum*d)
					boards = append(boards[:i], boards[i+1:]...)
					break
				}
			}
		}
	}
	fmt.Println(scores[0])
	fmt.Println(scores[len(scores)-1])
}

package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
	"unicode"
)

type State struct {
	Pos  image.Point
	Keys [26]bool
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	maze := map[image.Point]rune{}
	var start State
	var goal [26]bool

	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '@' {
				start = State{Pos: image.Point{x, y}}
				r = '.'
			}
			if unicode.IsLower(r) {
				goal[r-'a'] = true
			}
			maze[image.Point{x, y}] = r
		}
	}

	dist := map[State]int{start: 0}
	queue := []State{start}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.Keys == goal {
			fmt.Println(dist[state])
			return
		}

		for _, d := range []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			nextState := State{state.Pos.Add(d), state.Keys}
			nextTile := maze[nextState.Pos]

			if nextTile == '#' || unicode.IsUpper(nextTile) && !state.Keys[nextTile-'A'] {
				continue
			}
			if unicode.IsLower(nextTile) {
				nextState.Keys[nextTile-'a'] = true
			}
			if _, ok := dist[nextState]; !ok {
				dist[nextState] = dist[state] + 1
				queue = append(queue, nextState)
			}
		}
	}
}

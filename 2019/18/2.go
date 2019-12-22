package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
	"unicode"
)

type State struct {
	Pos    [4]image.Point
	Keys   [26]bool
	Active int
}

func main() {
	input, _ := ioutil.ReadFile("input2.txt")
	maze := map[image.Point]rune{}
	var start State
	var goal [26]bool
	robots := 0

	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '@' {
				start.Pos[robots] = image.Point{x, y}
				r = '.'
				robots++
			}
			if unicode.IsLower(r) {
				goal[r-'a'] = true
			}
			maze[image.Point{x, y}] = r
		}
	}

	dist := map[State]int{}
	queue := []State{}

	for i := 0; i < robots; i++ {
		start.Active = i
		dist[start] = 0
		queue = append(queue, start)
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.Keys == goal {
			fmt.Println(dist[state])
			return
		}

		for _, d := range []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			nextState := state
			nextState.Pos[state.Active] = state.Pos[state.Active].Add(d)
			nextTile := maze[nextState.Pos[state.Active]]

			if nextTile == '#' || unicode.IsUpper(nextTile) && !state.Keys[nextTile-'A'] {
				continue
			}
			if unicode.IsLower(nextTile) {
				nextState.Keys[nextTile-'a'] = true
			}
			for i := 0; i < robots; i++ {
				if i != state.Active && nextState.Keys == state.Keys {
					continue
				}
				nextState.Active = i
				if _, ok := dist[nextState]; !ok {
					dist[nextState] = dist[state] + 1
					queue = append(queue, nextState)
				}
			}
		}
	}
}

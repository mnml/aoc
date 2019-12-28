package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
	"unicode"
)

type State struct {
	P image.Point
	L int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	maze := map[image.Point]rune{}
	var w, h int

	for y, s := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		for x, r := range s {
			maze[image.Point{x, y}] = r
			w, h = x+1, y+1
		}
	}

	delta := []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	labels := map[string][]State{}
	portals := map[image.Point]State{}

	for p := range maze {
		if maze[p] != '.' {
			continue
		}

		for _, d := range delta {
			if !unicode.IsUpper(maze[p.Add(d)]) {
				continue
			}

			label := string(maze[p.Add(d)]) + string(maze[p.Add(d).Add(d)])
			if d.X < 0 || d.Y < 0 {
				label = string(maze[p.Add(d).Add(d)]) + string(maze[p.Add(d)])
			}

			labels[label] = append(labels[label], State{p, 0})
			if len(labels[label]) == 2 {
				portals[labels[label][0].P] = labels[label][1]
				portals[labels[label][1].P] = labels[label][0]
			}
		}
	}

	for p := range portals {
		portals[p] = State{portals[p].P, -1}
		if p.X > 2 && p.X < w-3 && p.Y > 2 && p.Y < h-3 {
			portals[p] = State{portals[p].P, 1}
		}
	}

	dist := map[State]int{labels["AA"][0]: 0}
	queue := []State{labels["AA"][0]}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state == labels["ZZ"][0] {
			fmt.Println(dist[state])
			return
		}

		for _, d := range delta {
			next := State{state.P.Add(d), state.L}
			if n, ok := portals[state.P]; ok && unicode.IsUpper(maze[next.P]) {
				next = State{n.P, state.L + n.L}
			}

			if _, ok := dist[next]; !ok && maze[next.P] == '.' && next.L >= 0 {
				dist[next] = dist[state] + 1
				queue = append(queue, next)
			}
		}
	}
}

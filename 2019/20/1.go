package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	maze := map[image.Point]rune{}

	for y, s := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		for x, r := range s {
			maze[image.Point{x, y}] = r
		}
	}

	delta := []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	labels := map[string][]image.Point{}
	portals := map[image.Point]image.Point{}

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

			labels[label] = append(labels[label], p)
			if len(labels[label]) >= 2 {
				portals[labels[label][0]] = labels[label][1]
				portals[labels[label][1]] = labels[label][0]
			}
		}
	}

	dist := map[image.Point]int{labels["AA"][0]: 0}
	queue := []image.Point{labels["AA"][0]}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		if point == labels["ZZ"][0] {
			fmt.Println(dist[point])
			return
		}

		for _, d := range delta {
			next := point.Add(d)
			if n, ok := portals[point]; ok && unicode.IsUpper(maze[next]) {
				next = n
			}

			if _, ok := dist[next]; !ok && maze[next] == '.' {
				dist[next] = dist[point] + 1
				queue = append(queue, next)
			}
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const isize = 12
const tsize = 10

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	tiles := map[int][]string{}
	edges := map[string][]int{}
	for _, s := range split {
		var id int
		fmt.Sscanf(s, "Tile %d:", &id)
		tiles[id] = strings.Split(s, "\n")[1:]
		for _, e := range getEdges(tiles[id]) {
			edges[canon(e)] = append(edges[canon(e)], id)
		}
	}

	counts := map[int]int{}
	var corner int

	part1 := 1
	for _, v := range edges {
		if len(v) == 1 {
			if counts[v[0]]++; counts[v[0]] == 2 {
				corner = v[0]
				part1 *= corner
			}
		}
	}
	fmt.Println(part1)

	order := [isize][isize][]string{{tiles[corner]}}
	for i := 1; len(edges[canon(order[0][0][0])]) != 1 || len(edges[canon(col(order[0][0], 0))]) != 1; i++ {
		order[0][0] = rotate(tiles[corner], i)
	}
	delete(tiles, corner)

	image := []string{}
	for y := 0; y < isize; y++ {
		for x := 0; x < isize; x++ {
			if y == 0 && x == 0 {
				continue
			}

			for id, tile := range tiles {
				if (y == 0 || hasEdge(tile, order[y-1][x][tsize-1])) && (x == 0 || hasEdge(tile, col(order[y][x-1], tsize-1))) {
					order[y][x] = tile
					for i := 1; y != 0 && order[y][x][0] != order[y-1][x][tsize-1] || x != 0 && col(order[y][x], 0) != col(order[y][x-1], tsize-1); i++ {
						order[y][x] = rotate(tile, i)
					}
					delete(tiles, id)
					break
				}
			}
		}
	}

	part2 := 0
	for y := 0; y < isize; y++ {
		for i := 1; i < tsize-1; i++ {
			image = append(image, "")
			for x := 0; x < isize; x++ {
				image[y*(tsize-2)+i-1] += order[y][x][i][1 : tsize-1]
				part2 += strings.Count(order[y][x][i][1:tsize-1], "#")
			}
		}
	}

	orig := make([]string, len(image))
	copy(orig, image)
	for i := 1; i <= 8; i++ {
		monster := "..................#.\n#....##....##....###\n.#..#..#..#..#..#..."
		split := strings.Split(monster, "\n")

		for y := 0; y < len(image)-len(split); y++ {
		search:
			for x := 0; x < len(image[0])-len(split[0]); x++ {
				for i, s := range split {
					if match, _ := regexp.MatchString(s, image[y+i][x:x+len(s)]); !match {
						continue search
					}
				}
				part2 -= strings.Count(monster, "#")
			}
		}
		image = rotate(orig, i)
	}
	fmt.Println(part2)
}

func getEdges(tile []string) []string {
	return []string{tile[0], col(tile, len(tile[0])-1), tile[len(tile)-1], col(tile, 0)}
}

func hasEdge(tile []string, edge string) bool {
	for _, e := range getEdges(tile) {
		if canon(e) == canon(edge) {
			return true
		}
	}
	return false
}

func reverse(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return
}

func canon(s string) string {
	if rs := reverse(s); rs < s {
		return rs
	}
	return s
}

func col(tile []string, i int) (col string) {
	for _, s := range tile {
		col += string(s[i])
	}
	return
}

func flip(tile []string) (flip []string) {
	for _, s := range tile {
		flip = append(flip, reverse(s))
	}
	return
}

func rotate(tile []string, n int) (rot []string) {
	if n%8 >= 4 {
		tile = flip(tile)
	}
	for ; n > 0; n-- {
		rot = []string{}
		for i := range tile {
			rot = append(rot, reverse(col(tile, i)))
		}
		tile = rot
	}
	return
}

package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"regexp"
	"strings"
)

type Tile []string

func (t Tile) Col(i int) (col string) {
	for _, s := range t {
		col += string(s[i])
	}
	return
}

func (t Tile) Edges() []string {
	return []string{t[0], t.Col(len(t[0]) - 1), t[len(t)-1], t.Col(0)}
}

func (t Tile) Rotations() []Tile {
	rot := make([]Tile, 8)
	for r := 0; r < 8; r += 2 {
		for _, s := range t {
			rot[r] = append(rot[r], reverse(s))
		}
		for i := range t {
			rot[r+1] = append(rot[r+1], reverse(t.Col(i)))
		}
		t = rot[r+1]
	}
	return rot
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	tiles := map[int]Tile{}
	counts := map[string]int{}
	for _, s := range split {
		var id int
		fmt.Sscanf(s, "Tile %d:", &id)
		tiles[id] = strings.Split(s, "\n")[1:]
		for _, e := range tiles[id].Edges() {
			counts[e]++
			counts[reverse(e)]++
		}
	}

	imageSize, tileSize := int(math.Sqrt(float64(len(tiles)))), 10
	img := make(Tile, imageSize*(tileSize-2))
	order := map[image.Point]Tile{}

	part1 := 1
	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
		findTile:
			for i, t := range tiles {
				for _, r := range t.Rotations() {
					if (y == 0 && counts[r[0]] == 1 ||
						y != 0 && r[0] == order[image.Point{x, y - 1}][tileSize-1]) &&
						(x == 0 && counts[r.Col(0)] == 1 ||
							x != 0 && r.Col(0) == order[image.Point{x - 1, y}].Col(tileSize-1)) {
						if (y == 0 || y == imageSize-1) && (x == 0 || x == imageSize-1) {
							part1 *= i
						}

						for i := 0; i < tileSize-2; i++ {
							img[(tileSize-2)*y+i] += r[i+1][1 : tileSize-1]
						}

						order[image.Point{x, y}] = r
						delete(tiles, i)
						break findTile
					}
				}
			}
		}
	}
	fmt.Println(part1)

	mon := []string{"..................#.", "#....##....##....###", ".#..#..#..#..#..#..."}
	nmon := 0
	for _, r := range img.Rotations() {
		for y := 0; y < len(r)-len(mon); y++ {
		findMon:
			for x := 0; x < len(r[0])-len(mon[0]); x++ {
				for i, s := range mon {
					if match, _ := regexp.MatchString(s, r[y+i][x:x+len(s)]); !match {
						continue findMon
					}
				}
				nmon++
			}
		}
	}
	fmt.Println(strings.Count(strings.Join(img, ""), "#") - nmon*strings.Count(strings.Join(mon, ""), "#"))
}

func reverse(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return
}

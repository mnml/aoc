package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

type Light struct {
	Pos image.Point
	Vel image.Point
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lights := []Light{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		l := Light{}
		fmt.Sscanf(s, "position=<%d, %d> velocity=<%d, %d>", &l.Pos.X, &l.Pos.Y, &l.Vel.X, &l.Vel.Y)
		lights = append(lights, l)
	}

	seconds := 0
	for {
		tl, br := image.Point{MaxInt, MaxInt}, image.Point{MinInt, MinInt}
		points := map[image.Point]struct{}{}
		for i := range lights {
			p := lights[i].Pos.Add(lights[i].Vel)
			lights[i].Pos = p
			tl, br = image.Point{Min(tl.X, p.X), Min(tl.Y, p.Y)}, image.Point{Max(br.X, p.X), Max(br.Y, p.Y)}
			points[p] = struct{}{}
		}
		seconds++
		if br.Y-tl.Y <= 9 {
			for y := tl.Y; y <= br.Y; y++ {
				for x := tl.X; x <= br.X; x++ {
					_, ok := points[image.Point{x, y}]
					fmt.Print(map[bool]string{true: "#", false: " "}[ok])
				}
				fmt.Println()
			}
			fmt.Println(seconds)
			break
		}
	}
}

func Max(x ...int) int {
	max := x[0]
	for _, v := range x[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(x ...int) int {
	min := x[0]
	for _, v := range x[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

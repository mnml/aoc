package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	tl, br := image.Point{MaxInt, MaxInt}, image.Point{MinInt, MinInt}
	points := map[image.Point]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var x, y int
		fmt.Sscanf(s, "%d, %d", &x, &y)
		points[image.Point{x, y}] = 0
		tl, br = image.Point{Min(tl.X, x), Min(tl.Y, y)}, image.Point{Max(br.X, x), Max(br.Y, y)}
	}

	area := 0
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			mp, md := image.Point{MaxInt, MaxInt}, MaxInt
			sum := 0

			for p := range points {
				d := Abs(x-p.X) + Abs(y-p.Y)
				if d < md {
					mp, md = p, d
				} else if d == md {
					mp = image.Point{MinInt, MinInt}
				}
				sum += d
			}

			if a, ok := points[mp]; ok && a != -1 {
				points[mp]++
				if x <= tl.X || y <= tl.Y || x >= br.X || y >= br.Y {
					points[mp] = -1
				}
			}

			if sum < 10000 {
				area++
			}
		}
	}

	max := 0
	for _, v := range points {
		max = Max(v, max)
	}
	fmt.Println(max)
	fmt.Println(area)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

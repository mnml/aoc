package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	asteroids := []image.Point{}
	bucket := map[float64][]image.Point{}
	station := image.Point{}

	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '#' {
				asteroids = append(asteroids, image.Point{x, y})
			}
		}
	}

	for _, s := range asteroids {
		b := map[float64][]image.Point{}
		for _, a := range asteroids {
			if a != s {
				angle := -math.Atan2(float64(a.X-s.X), float64(a.Y-s.Y))
				b[angle] = append(b[angle], a)
			}
		}
		if len(b) > len(bucket) {
			bucket = b
			station = s
		}
	}
	fmt.Println(len(bucket))

	for a := range bucket {
		sort.Slice(bucket[a], func(i, j int) bool {
			return dist(station, bucket[a][i]) > dist(station, bucket[a][j])
		})
		for len(bucket[a]) > 1 {
			i := a + 2*math.Pi*float64(len(bucket[a])-1)
			bucket[i] = append(bucket[i], bucket[a][0])
			bucket[a] = bucket[a][1:]
		}
	}

	angles := []float64{}
	for a := range bucket {
		angles = append(angles, a)
	}
	sort.Float64s(angles)
	fmt.Println(100*bucket[angles[199]][0].X + bucket[angles[199]][0].Y)
}

func dist(a, b image.Point) float64 {
	return math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2))
}

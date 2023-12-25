package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y, Z float64
}

func main() {
	input, _ := os.ReadFile("input.txt")

	hail := [][2]Point{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var p, v Point
		fmt.Sscanf(s, "%f, %f, %f @ %f, %f, %f", &p.X, &p.Y, &p.Z, &v.X, &v.Y, &v.Z)
		hail = append(hail, [2]Point{p, v})
	}

	part1 := 0
	for i, h1 := range hail {
		for _, h2 := range hail[i+1:] {
			p1, v1, p2, v2 := h1[0], h1[1], h2[0], h2[1]
			det := v2.X*v1.Y - v2.Y*v1.X
			t1 := ((p2.Y-p1.Y)*v2.X - (p2.X-p1.X)*v2.Y) / det
			t2 := ((p2.Y-p1.Y)*v1.X - (p2.X-p1.X)*v1.Y) / det
			if x, y := p1.X+t1*v1.X, p1.Y+t1*v1.Y; t1 > 0 && t2 > 0 &&
				x >= 200000000000000 && x <= 400000000000000 &&
				y >= 200000000000000 && y <= 400000000000000 {
				part1++
			}
		}
	}
	fmt.Println(part1)

	// https://github.com/0e4ef622/aoc/blob/master/2023/day24/a.py
	solve := func(h1, h2, h3 [2]Point) float64 {
		x1, y1, z1, vx1, vy1, vz1 := h1[0].X, h1[0].Y, h1[0].Z, h1[1].X, h1[1].Y, h1[1].Z
		x2, y2, z2, vx2, vy2, vz2 := h2[0].X, h2[0].Y, h2[0].Z, h2[1].X, h2[1].Y, h2[1].Z
		x3, y3, z3, vx3, vy3, vz3 := h3[0].X, h3[0].Y, h3[0].Z, h3[1].X, h3[1].Y, h3[1].Z
		yz := y1*(z2-z3) + y2*(-z1+z3) + y3*(z1-z2)
		xz := x1*(-z2+z3) + x2*(z1-z3) + x3*(-z1+z2)
		xy := x1*(y2-y3) + x2*(-y1+y3) + x3*(y1-y2)
		vxvy := vx1*(vy2-vy3) + vx2*(-vy1+vy3) + vx3*(vy1-vy2)
		vxvz := vx1*(-vz2+vz3) + vx2*(vz1-vz3) + vx3*(-vz1+vz2)
		vyvz := vy1*(vz2-vz3) + vy2*(-vz1+vz3) + vy3*(vz1-vz2)
		n := (vx2-vx3)*yz + (vy2-vy3)*xz + (vz2-vz3)*xy
		d := (z2-z3)*vxvy + (y2-y3)*vxvz + (x2-x3)*vyvz
		return n / d
	}

	t1 := solve(hail[0], hail[1], hail[2])
	t2 := solve(hail[1], hail[0], hail[2])

	p1, v1, p2, v2 := hail[0][0], hail[0][1], hail[1][0], hail[1][1]
	c1 := Point{p1.X + t1*v1.X, p1.Y + t1*v1.Y, p1.Z + t1*v1.Z}
	c2 := Point{p2.X + t2*v2.X, p2.Y + t2*v2.Y, p2.Z + t2*v2.Z}
	v := Point{(c2.X - c1.X) / (t2 - t1), (c2.Y - c1.Y) / (t2 - t1), (c2.Z - c1.Z) / (t2 - t1)}
	p := Point{p1.X + v1.X*t1 - v.X*t1, p1.Y + v1.Y*t1 - v.Y*t1, p1.Z + v1.Z*t1 - v.Z*t1}

	fmt.Println(int(math.Round(p.X + p.Y + p.Z)))
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Moon struct {
	P [3]int
	V [3]int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	moons := []Moon{}

	for i, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		moons = append(moons, Moon{})
		fmt.Sscanf(s, "<x=%d, y=%d, z=%d>", &moons[i].P[0], &moons[i].P[1], &moons[i].P[2])
	}

	init := append([]Moon(nil), moons...)
	period := make([]int, 3)
	for s := 1; period[0] == 0 || period[1] == 0 || period[2] == 0; s++ {
		for i := range period {
			for m1 := range moons {
				for m2 := range moons {
					if moons[m1].P[i] < moons[m2].P[i] {
						moons[m1].V[i]++
					} else if moons[m1].P[i] > moons[m2].P[i] {
						moons[m1].V[i]--
					}
				}
			}

			for m := range moons {
				moons[m].P[i] += moons[m].V[i]
			}

			if period[i] == 0 {
				for m := range moons {
					if moons[m].V[i] != init[m].V[i] {
						break
					}
					if m == len(moons)-1 {
						period[i] = s * 2
					}
				}
			}
		}
	}

	fmt.Println(lcm(period...))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(x ...int) int {
	if len(x) == 0 {
		return 0
	} else if len(x) == 2 {
		for x[1] != 0 {
			x[0], x[1] = x[1], x[0]%x[1]
		}
	} else if len(x) > 2 {
		return gcd(x[0], gcd(x[1:]...))
	}
	return abs(x[0])
}

func lcm(x ...int) int {
	if len(x) > 2 {
		return lcm(x[0], lcm(x[1:]...))
	} else if x[0] == 0 && x[1] == 0 {
		return 0
	}
	return abs(x[0]*x[1]) / gcd(x[0], x[1])
}

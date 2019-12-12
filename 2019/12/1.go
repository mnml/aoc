package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	for s := 0; s < 1000; s++ {
		for i := 0; i < 3; i++ {
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
		}
	}

	energy := 0
	for m := range moons {
		for i := range moons[m].P {
			for j := range moons[m].V {
				energy += int(math.Abs(float64(moons[m].P[i] * moons[m].V[j])))
			}
		}
	}

	fmt.Println(energy)
}

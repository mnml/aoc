package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	var seeds []int
	json.Unmarshal([]byte("["+strings.Join(strings.Fields(strings.Split(split[0], ": ")[1]), ",")+"]"), &seeds)

	maps := [][][3]int{}
	for i, s := range split[1:] {
		maps = append(maps, [][3]int{})

		for j, s := range strings.Split(strings.Split(s, ":\n")[1], "\n") {
			maps[i] = append(maps[i], [3]int{})
			fmt.Sscanf(s, "%d %d %d", &maps[i][j][0], &maps[i][j][1], &maps[i][j][2])
		}
	}

	calc := func(seed int) int {
		for _, m := range maps {
			for _, r := range m {
				if seed >= r[1] && seed < r[1]+r[2] {
					seed = r[0] + seed - r[1]
					break
				}
			}
		}
		return seed
	}

	part1, part2 := math.MaxInt, math.MaxInt
	for i, s := range seeds {
		part1 = slices.Min([]int{part1, calc(s)})
		if i%2 == 0 {
			for s := seeds[i]; s < seeds[i]+seeds[i+1]; s++ {
				part2 = slices.Min([]int{part2, calc(s)})
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

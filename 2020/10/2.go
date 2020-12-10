package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	jolts := make([]int, len(split))
	for i, s := range split {
		jolts[i], _ = strconv.Atoi(s)
	}
	sort.Ints(jolts)
	jolts = append([]int{0}, append(jolts, jolts[len(jolts)-1]+3)...)

	diff, memo := map[int]int{}, map[int]int{0: 1}
	for i, v := range jolts[1:] {
		diff[v-jolts[i]]++
		memo[v] = memo[v-1] + memo[v-2] + memo[v-3]
	}
	fmt.Println(diff[1] * diff[3])
	fmt.Println(memo[jolts[len(jolts)-1]])
}

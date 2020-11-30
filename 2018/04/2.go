package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	sort.Strings(split)

	re := regexp.MustCompile(`.*:(\d+)\] (\w+) .(\d+)?`)
	records := map[int]map[int]int{}
	var id, falls, wakes int
	var maxID, maxMin int

	for _, s := range split {
		m := re.FindStringSubmatch(s)
		switch m[2] {
		case "Guard":
			id, _ = strconv.Atoi(m[3])
			if _, ok := records[id]; !ok {
				records[id] = map[int]int{}
			}
		case "falls":
			falls, _ = strconv.Atoi(m[1])
		case "wakes":
			wakes, _ = strconv.Atoi(m[1])
			for min := falls; min < wakes; min++ {
				records[id][min]++
				if records[id][min] > records[maxID][maxMin] {
					maxID, maxMin = id, min
				}
			}
		}
	}

	fmt.Println(maxID * maxMin)
}

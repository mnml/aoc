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

	ids := []int{}
	for _, s := range strings.Fields(string(input)) {
		bin := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(s)
		id, _ := strconv.ParseUint(bin, 2, 10)
		ids = append(ids, int(id))
	}
	sort.Ints(ids)

	fmt.Println(ids[len(ids)-1])
	for i := range ids {
		if ids[i+1] != ids[i]+1 {
			fmt.Println(ids[i] + 1)
			break
		}
	}
}

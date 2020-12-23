package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	labels := strings.TrimSpace(string(input))

	fmt.Println(run(labels, len(labels), 100))
	fmt.Println(run(labels, 1000000, 10000000))
}

func run(labels string, ncups int, moves int) (ans int) {
	cups := ring.New(ncups)
	ps := map[int]*ring.Ring{}

	for i := 1; i <= ncups; i++ {
		if cups.Value = i; i <= len(labels) {
			cups.Value = int(labels[i-1] - '0')
		}
		ps[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	for i := 0; i < moves; i++ {
		pick := cups.Unlink(3)
		dest := (ncups+cups.Value.(int)-2)%ncups + 1

		unavail := map[int]bool{}
		for i := 0; i < 3; i++ {
			unavail[pick.Value.(int)] = true
			pick = pick.Next()
		}

		for unavail[dest] {
			dest = (ncups+dest-2)%ncups + 1
		}

		ps[dest].Link(pick)
		cups = cups.Next()
	}

	if ncups > len(labels) {
		return ps[1].Next().Value.(int) * ps[1].Move(2).Value.(int)
	}
	ps[1].Unlink(len(labels) - 1).Do(func(p interface{}) {
		ans = ans*10 + p.(int)
	})
	return
}

package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	var nplayers, nmarbles int
	fmt.Sscanf(string(input), "%d players; last marble is worth %d points", &nplayers, &nmarbles)
	marbles := ring.New(1)
	marbles.Value = 0
	players := make([]int, nplayers)

	max := 0
	for i := 1; i <= nmarbles*100; i++ {
		if i%23 == 0 {
			marbles = marbles.Move(-6)
			players[i%len(players)] += i + marbles.Move(-2).Link(marbles).Value.(int)
			if players[i%len(players)] > max {
				max = players[i%len(players)]
			}
		} else {
			marbles = marbles.Next().Link(ring.New(1)).Prev()
			marbles.Value = i
		}
		if i == nmarbles {
			fmt.Println(max)
		}
	}
	fmt.Println(max)
}

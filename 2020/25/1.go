package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	var card, door int
	fmt.Sscanf(string(input), "%d\n%d", &card, &door)

	loop := 0
	for k := 1; k != card; loop++ {
		k = k * 7 % 20201227
	}

	key := 1
	for l := 0; l < loop; l++ {
		key = key * door % 20201227
	}
	fmt.Println(key)
}

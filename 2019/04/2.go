package main

import (
	"fmt"
	"strconv"
)

func main() {
	valid := 0

passLoop:
	for i := 240298; i <= 784956; i++ {
		str := strconv.Itoa(i)
		count := map[byte]int{str[0]: 1}

		for i := 1; i < len(str); i++ {
			if str[i-1] > str[i] {
				continue passLoop
			}
			count[str[i]]++
		}

		for _, v := range count {
			if v == 2 {
				valid++
				break
			}
		}
	}

	fmt.Println(valid)
}

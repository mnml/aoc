package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	rack := map[string][]string{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		split := strings.Fields(s)
		rack[split[0][:3]] = split[1:]
	}

	var paths func(string, bool, bool) int
	cache := map[string]int{}

	paths = func(dev string, dac, fft bool) (n int) {
		if n, ok := cache[fmt.Sprint(dev, dac, fft)]; ok {
			return n
		}
		defer func() { cache[fmt.Sprint(dev, dac, fft)] = n }()

		if dev == "out" && dac && fft {
			return 1
		}
		dac = dac || dev == "dac"
		fft = fft || dev == "fft"
		for _, d := range rack[dev] {
			n += paths(d, dac, fft)
		}
		return n
	}

	fmt.Println(paths("you", true, true))
	fmt.Println(paths("svr", false, false))
}

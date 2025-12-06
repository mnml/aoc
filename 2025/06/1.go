package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	lines := split[:len(split)-1]
	ops := strings.Fields(split[len(split)-1])

	nums1 := [][]int{}
	for x := range strings.Fields(lines[0]) {
		s := ""
		for y := range lines {
			s += strings.Fields(lines[y])[x] + " "
		}
		nums1 = append(nums1, parse(s))
	}

	transposed := ""
	for x := range lines[0] {
		s := ""
		for y := range lines {
			s += string(lines[y][x])
		}
		transposed += s + " "
		if strings.TrimSpace(s) == "" {
			transposed += "\n"
		}
	}

	nums2 := [][]int{}
	for _, line := range strings.Split(transposed, "\n") {
		nums2 = append(nums2, parse(line))
	}

	fmt.Println(calc(nums1, ops))
	fmt.Println(calc(nums2, ops))
}

func parse(s string) []int {
	ns := []int{}
	for _, s := range strings.Fields(s) {
		n, _ := strconv.Atoi(s)
		ns = append(ns, n)
	}
	return ns
}

func calc(nums [][]int, ops []string) int {
	result := 0
	for i, op := range ops {
		acc := nums[i][0]
		for _, n := range nums[i][1:] {
			switch op {
			case "+":
				acc += n
			case "*":
				acc *= n
			}
		}
		result += acc
	}
	return result
}

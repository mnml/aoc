package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	Children []Node
	Metadata []int
}

func process(license []int) (Node, []int) {
	n, nc, nm := Node{}, license[0], license[1]
	license = license[2:]
	for i := 0; i < nc; i++ {
		var c Node
		c, license = process(license)
		n.Children = append(n.Children, c)
	}
	n.Metadata = license[:nm]
	return n, license[nm:]
}

func sum(tree Node) (s int) {
	for _, m := range tree.Metadata {
		s += m
	}
	for _, c := range tree.Children {
		s += sum(c)
	}
	return
}

func value(tree Node) (v int) {
	if len(tree.Children) == 0 {
		v = sum(tree)
	} else {
		for _, m := range tree.Metadata {
			if i := m - 1; i >= 0 && i < len(tree.Children) {
				v += value(tree.Children[i])
			}
		}
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	license := []int{}

	for _, s := range strings.Fields(string(input)) {
		i, _ := strconv.Atoi(s)
		license = append(license, i)
	}

	tree, _ := process(license)
	fmt.Println(sum(tree))
	fmt.Println(value(tree))
}

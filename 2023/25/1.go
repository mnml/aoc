package main

import (
	"fmt"
	"maps"
	"math/rand"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	edges := []map[string]struct{}{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		split := strings.Split(s, ": ")
		for _, s := range strings.Fields(split[1]) {
			edges = append(edges, map[string]struct{}{split[0]: {}, s: {}})
		}
	}

	for {
		es := []map[string]struct{}{}
		for _, e := range edges {
			es = append(es, maps.Clone(e))
		}
		vs := func() int {
			vs := map[string]struct{}{}
			for _, e := range es {
				maps.Copy(vs, e)
			}
			return len(vs)
		}

		for vs() > 2 {
			edge := es[rand.Intn(len(es))]

			es = slices.DeleteFunc(es, func(e map[string]struct{}) bool {
				return maps.Equal(e, edge)
			})

			name := ""
			for v := range edge {
				name += v + " "
			}

			for e := range es {
				for v := range es[e] {
					if _, ok := edge[v]; ok {
						delete(es[e], v)
						es[e][name] = struct{}{}
					}
				}
			}
		}

		if len(es) != 3 {
			continue
		}

		part1 := 1
		for v := range es[0] {
			part1 *= len(strings.Fields(v))
		}
		fmt.Println(part1)
		break
	}
}

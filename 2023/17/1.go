package main

import (
	"container/heap"
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

type State struct {
	Pos image.Point
	Dir image.Point
}

func main() {
	input, _ := os.ReadFile("input.txt")

	grid, end := map[image.Point]int{}, image.Point{0, 0}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = int(r - '0')
			end = image.Point{x, y}
		}
	}

	run := func(min, max int) int {
		queue, seen := PQ[State]{}, map[State]struct{}{}
		queue.GPush(State{image.Point{0, 0}, image.Point{1, 0}}, 0)
		queue.GPush(State{image.Point{0, 0}, image.Point{0, 1}}, 0)

		for len(queue) > 0 {
			state, heat := queue.GPop()

			if state.Pos == end {
				return heat
			}
			if _, ok := seen[state]; ok {
				continue
			}
			seen[state] = struct{}{}

			for i := -max; i <= max; i++ {
				n := state.Pos.Add(state.Dir.Mul(i))
				if _, ok := grid[n]; !ok || i > -min && i < min {
					continue
				}
				h, s := 0, int(math.Copysign(1, float64(i)))
				for j := s; j != i+s; j += s {
					h += grid[state.Pos.Add(state.Dir.Mul(j))]
				}
				queue.GPush(State{n, image.Point{state.Dir.Y, state.Dir.X}}, heat+h)
			}
		}
		return -1
	}

	fmt.Println(run(1, 3))
	fmt.Println(run(4, 10))
}

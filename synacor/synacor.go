package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

const (
	size = 1 << 15
	nreg = 8
)

func main() {
	f, _ := os.Open("challenge.bin")
	fi, _ := f.Stat()

	mem := make([]uint16, size+nreg+1)
	binary.Read(f, binary.LittleEndian, mem[:fi.Size()/2])

	// coins := map[int]string{9: "blue", 2: "red", 5: "shiny", 7: "concave", 3: "corroded"}
	// for _, p := range perms(2, 3, 5, 7, 9) {
	// 	if p[0]+p[1]*p[2]*p[2]+p[3]*p[3]*p[3]-p[4] == 399 {
	// 		for _, v := range p {
	// 			fmt.Printf("use %s coin\n", coins[v])
	// 		}
	// 		break
	// 	}
	// }

	// disasm(mem[:fi.Size()/2])

	// for r7 := uint16(0); r7 < size; r7++ {
	// 	if telecheck(4, 1, r7, &[5][size]uint16{}) == 6 {
	// 		fmt.Println(r7)
	// 		break
	// 	}
	// }

	// fmt.Print(vault())

	// fmt.Println(putMyThingDownFlipItAndReverseIt("dxTdMYpudWbp"))

	mem[0x154d] = 0x178f
	mem[0x178b] = 1     // set
	mem[0x178c] = 32768 // r0
	mem[0x178d] = 6
	mem[0x178e] = 18    // ret
	mem[0x178f] = 1     // set
	mem[0x1790] = 32775 // r7
	mem[0x1791] = 25734 // (teleport check)
	mem[0x1792] = 6     // jmp
	mem[0x1793] = 0x15e5
	run(mem)
}

func run(mem []uint16) {
	pc := &mem[size+nreg]
	op := func() *uint16 {
		defer func() { *pc++ }()
		if mem[*pc] >= size {
			return &mem[mem[*pc]]
		}
		return &mem[*pc]
	}
	btou := func(b bool, t, f uint16) uint16 {
		if b {
			return t
		}
		return f
	}

	for {
		switch *op() {
		case 0:
			return
		case 1:
			*op() = *op()
		case 2:
			mem = append(mem, *op())
		case 3:
			*op(), mem = mem[len(mem)-1], mem[:len(mem)-1]
		case 4:
			*op() = btou(*op() == *op(), 1, 0)
		case 5:
			*op() = btou(*op() > *op(), 1, 0)
		case 6:
			*pc = *op()
		case 7:
			*pc = btou(*op() != 0, *op(), *pc)
		case 8:
			*pc = btou(*op() == 0, *op(), *pc)
		case 9:
			*op() = (*op() + *op()) % size
		case 10:
			*op() = *op() * *op() % size
		case 11:
			*op() = *op() % *op()
		case 12:
			*op() = *op() & *op()
		case 13:
			*op() = *op() | *op()
		case 14:
			*op() = ^*op() % size
		case 15:
			*op() = mem[*op()]
		case 16:
			mem[*op()] = *op()
		case 17:
			*pc, mem = *op(), append(mem, *pc)
		case 18:
			*pc, mem = mem[len(mem)-1], mem[:len(mem)-1]
		case 19:
			os.Stdout.Write([]byte{byte(*op())})
		case 20:
			b := make([]byte, 1)
			os.Stdin.Read(b)
			*op() = uint16(b[0])
		}
	}
}

func perms(xs ...int) (ps [][]int) {
	if len(xs) == 1 {
		return [][]int{xs}
	}
	for i := range xs {
		c := make([]int, len(xs))
		copy(c, xs)
		for _, x := range perms(append(c[:i], c[i+1:]...)...) {
			ps = append(ps, append([]int{xs[i]}, x...))
		}
	}
	return
}

func disasm(mem []uint16) {
	names := []string{"halt", "set", "push", "pop", "eq", "gt", "jmp", "jt", "jf", "add",
		"mult", "mod", "and", "or", "not", "rmem", "wmem", "call", "ret", "out", "in", "noop"}
	nops := map[uint16]int{0: 0, 1: 2, 2: 1, 3: 1, 4: 3, 5: 3, 6: 1, 7: 2, 8: 2, 9: 3, 10: 3,
		11: 3, 12: 3, 13: 3, 14: 2, 15: 2, 16: 2, 17: 1, 18: 0, 19: 1, 20: 1, 21: 0}

	for pc := 0; pc < len(mem); pc += nops[mem[pc]] + 1 {
		if _, ok := nops[mem[pc]]; !ok {
			continue
		}

		fmt.Printf("%04x %s", pc, names[mem[pc]])
		for i := 1; i < nops[mem[pc]]+1; i++ {
			fmt.Printf(map[bool]string{true: " r%d", false: " %#x"}[mem[pc+i] >= size], mem[pc+i]%size)
		}
		fmt.Println()
	}
}

func telecheck(r0, r1, r7 uint16, cache *[5][size]uint16) (r uint16) {
	if cache[r0][r1] != 0 {
		return cache[r0][r1]
	}
	defer func() { r %= size; cache[r0][r1] = r }()

	switch r0 {
	case 2:
		return (r1+2)*r7 + r1 + 1
	case 1:
		return r7 + r1 + 1
	case 0:
		return r1 + 1
	}

	if r1 == 0 {
		return telecheck(r0-1, r7, r7, cache)
	}
	return telecheck(r0-1, telecheck(r0, r1-1, r7, cache), r7, cache)
}

func vault() string {
	type State struct {
		Pos image.Point
		Val int
	}

	s := `*  8  -  1
	      4  *  11 *
	      +  4  -  18
	      22  - 9  *`

	grid := map[image.Point]string{}
	for y, s := range strings.Split(s, "\n") {
		for x, s := range strings.Fields(s) {
			grid[image.Point{x, y}] = s
		}
	}

	start := State{image.Point{0, 3}, 22}
	end := State{image.Point{3, 0}, 30}

	queue := []State{start}
	path := map[State]string{start: ""}

	for len(queue) > 0 {
		if queue[0] == end {
			return path[queue[0]]
		}

		for s, d := range map[string]image.Point{"north": {0, -1}, "east": {1, 0}, "south": {0, 1}, "west": {-1, 0}} {
			next := State{queue[0].Pos.Add(d), queue[0].Val}

			switch v, _ := strconv.Atoi(grid[next.Pos]); grid[queue[0].Pos] {
			case "*":
				next.Val *= v
			case "+":
				next.Val += v
			case "-":
				next.Val -= v
			}

			if _, ok := grid[next.Pos]; !ok ||
				next.Pos == start.Pos ||
				next.Pos == end.Pos && next.Val != end.Val {
				continue
			}

			if _, ok := path[next]; !ok {
				path[next] = path[queue[0]] + s + "\n"
				queue = append(queue, next)
			}
		}

		queue = queue[1:]
	}

	return ""
}

func putMyThingDownFlipItAndReverseIt(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return strings.NewReplacer("b", "d", "d", "b", "p", "q", "q", "p").Replace(rs)
}

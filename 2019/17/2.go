package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	mem := map[int]int{}
	for i, s := range strings.Split(strings.TrimSpace(string(input)), ",") {
		mem[i], _ = strconv.Atoi(s)
	}

	cur := image.Point{0, 0}
	scaff := map[image.Point]rune{}
	var pos image.Point
	var dir int

	for _, v := range run(mem, "") {
		switch r := rune(v); r {
		case '\n':
			cur.X = 0
			cur.Y++
		case '^', '<', 'v', '>':
			pos = cur
			dir = map[rune]int{'^': 0, '<': 1, 'v': 2, '>': 3}[r]
			r = '#'
			fallthrough
		default:
			scaff[cur] = r
			cur.X++
		}
	}

	d := []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	sum := 0
	for p := range scaff {
		if scaff[p.Add(d[0])] == '#' && scaff[p.Add(d[1])] == '#' &&
			scaff[p.Add(d[2])] == '#' && scaff[p.Add(d[3])] == '#' {
			sum += p.X * p.Y
		}
	}
	fmt.Println(sum)

	path := ""
	for {
		if scaff[pos.Add(d[(dir+1)%4])] == '#' {
			dir = (dir + 1) % 4
			path += "L,"
		} else if scaff[pos.Add(d[(dir+3)%4])] == '#' {
			dir = (dir + 3) % 4
			path += "R,"
		} else {
			break
		}
		dist := 0
		for scaff[pos.Add(d[dir])] == '#' {
			pos = pos.Add(d[dir])
			dist++
		}
		path += strconv.Itoa(dist) + ","
	}

	var a, b, c string
compress:
	for i := 2; i <= 21; i++ {
		for j := 2; j <= 21; j++ {
			for k := 2; k <= 21; k++ {
				str := path
				a = str[:i]
				str = strings.ReplaceAll(str, a, "")
				b = str[:j]
				str = strings.ReplaceAll(str, b, "")
				c = str[:k]
				str = strings.ReplaceAll(str, c, "")
				if str == "" {
					break compress
				}
			}
		}
	}

	a = strings.Trim(a, ",")
	b = strings.Trim(b, ",")
	c = strings.Trim(c, ",")
	path = strings.ReplaceAll(path, a, "A")
	path = strings.ReplaceAll(path, b, "B")
	path = strings.ReplaceAll(path, c, "C")
	path = strings.Trim(path, ",")

	mem[0] = 2
	out := run(mem, strings.Join([]string{path, a, b, c, "n", ""}, "\n"))
	fmt.Println(out[len(out)-1])
}

func run(init map[int]int, in string) (out []int) {
	mem := map[int]int{}
	for i, v := range init {
		mem[i] = v
	}

	ip := 0
	rb := 0

	for {
		ins := fmt.Sprintf("%05d", mem[ip])
		op, _ := strconv.Atoi(ins[3:])
		par := func(i int) (addr int) {
			switch ins[3-i] {
			case '1':
				return ip + i
			case '2':
				return rb + mem[ip+i]
			default:
				return mem[ip+i]
			}
		}

		switch op {
		case 1:
			mem[par(3)] = mem[par(1)] + mem[par(2)]
		case 2:
			mem[par(3)] = mem[par(1)] * mem[par(2)]
		case 3:
			mem[par(1)] = int(in[0])
			in = in[1:]
		case 4:
			out = append(out, mem[par(1)])
		case 5:
			if mem[par(1)] != 0 {
				ip = mem[par(2)]
				continue
			}
		case 6:
			if mem[par(1)] == 0 {
				ip = mem[par(2)]
				continue
			}
		case 7:
			if mem[par(1)] < mem[par(2)] {
				mem[par(3)] = 1
			} else {
				mem[par(3)] = 0
			}
		case 8:
			if mem[par(1)] == mem[par(2)] {
				mem[par(3)] = 1
			} else {
				mem[par(3)] = 0
			}
		case 9:
			rb += mem[par(1)]
		case 99:
			return
		}

		ip += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var intslice []int

func main() {
	data()
	part1()
	part2()
}

func part2() int {
	d := slices.Clone(intslice)
	p := slices.Clone(intslice)
loop1:
	for {
		for k, i := range d {
			// fmt.Println("start:", k, "i:", i, "d[0]:", d[0], "p[0]:", p[0], "total:", i+d[0]+p[0])
			if i+d[0]+p[0] == 2020 {
				ok, _ := fmt.Println("Part2 :", i*d[0]*p[0], i, d[0], p[0])
				return ok
			} else if k == len(d)-1 {
				p = slices.Delete(p, 0, 1)
				// fmt.Println("first statement")
				continue loop1
			} else if len(p) == 1 {
				d = slices.Delete(d, 0, 1)
				p = slices.Clone(intslice)
				// fmt.Println("second statement")
				continue loop1
			}
			// fmt.Println("end:", k, "\n", "i:", i, "d[0]:", d[0], "p[0]:", p[0], "total:", i+d[0]+p[0])
		}
		fmt.Println("stop")
		break
	}
	return 0
}

func part1() int {
	d := slices.Clone(intslice)
loop:
	for {
		for k, i := range d {
			if i+d[0] == 2020 {
				ok, _ := fmt.Println("Part1 :", i*d[0], i, d[0])
				return ok
			} else if k == len(d)-1 {
				d = slices.Delete(d, 0, 1)
				continue loop
			}
		}
		break
	}
	return 0
}

func data() []int {
	data, _ := os.ReadFile("data/day01.txt")
	line := bytes.Split(data, []byte("\n"))
	for _, num := range line {
		i, _ := strconv.Atoi(string(num))
		if i == 0 {
		} else {
			intslice = append(intslice, i)
		}
	}
	return intslice
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var list_one []int
var list_two []int

func main() {
	data()
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func part1() int {
	var answer int
	a := slices.Clone(list_one)
	b := slices.Clone(list_two)
	slices.Sort(a)
	slices.Sort(b)

	for k, _ := range list_one {
		a := a[k] - b[k]
		answer = answer + abs(a)
	}
	return answer
}
func part2() int {
	var answer int

	a := slices.Clone(list_one)
	b := slices.Clone(list_two)

	for k, _ := range a {
		pos := k
		amt := 0
		for _, l := range b {
			if a[pos] == l {
				amt = amt + 1
			} else {
				continue
			}
		}
		answer = answer + (a[pos] * amt)
	}

	return answer
}
func data() {
	data, _ := os.ReadFile("data/day01.txt")
	line := bytes.SplitN(data, []byte("\n"), (len(string(data)) - 6))
	for _, k := range line {
		split := bytes.Join(bytes.Fields(k), []byte(" "))
		t, d, _ := bytes.Cut(split, []byte(" "))

		conv1, _ := strconv.Atoi(string(t))
		conv2, _ := strconv.Atoi(string(d))
		list_one = append(list_one, conv1)
		list_two = append(list_two, conv2)
		// fmt.Println("1", split, "2", t, "3", d, "4", conv1, "5", conv2)
	}
	list_one = list_one[0 : len(list_one)-1]
	list_two = list_two[0 : len(list_two)-1]
	return
}
func abs(i int) int {
	if i >= 0 {
		return i * 1
	} else {
		return i * -1
	}
}

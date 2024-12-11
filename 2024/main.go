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
}

func part1() int {
	var answer int
	slices.Sort(list_one)
	slices.Sort(list_two)

	for k, _ := range list_one {
		a := list_one[k] - list_two[k]
		answer = answer + abs(a)
	}
	return answer
}
func part2() {

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

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("data ", data())
	fmt.Println("Part 1: ", part1(data()))
}
func part1(iss [][]int) int {
	var report_part []int
	var report [][]int
	for _, m := range iss {
		for k, _ := range m {
			if k == len(m)-1 {
				break
			} else {
				la := m[:len(m)-1]
				ls := m[1:]
				a := ls[k] - la[k]
				report_part = append(report_part, a)
			}
		}
		report = append(report, report_part)
		report_part = nil
	}
	answer := report_test(report)
	return answer
}
func report_test(iss [][]int) int {
	var success int
	fmt.Println(iss)

	for o, m := range iss {

		for k, l := range m {
			if abs(sum(m)) != asum(m) {
				// fmt.Println(m, sum(m), asum(m), abs(sum(m)), l)
				break
			} else if abs(l) > 3 {
				// fmt.Println("first", m, l)
				break
			} else if l == 0 {
				// fmt.Println("second", m, l)
				break
			}
			if k+1 == len(m) {
				fmt.Println(o, m, sum(m), asum(m), abs(sum(m)), l)
				success = success + 1
			}

		}

	}
	return success
}
func data() [][]int {
	data, _ := os.ReadFile("data/day02.txt")
	line := bytes.Split(data, []byte("\n"))

	var list [][]int

	for _, l := range line {
		var list_part []int
		a := bytes.Split(l, []byte(" "))
		for _, l := range a {
			i, _ := strconv.Atoi(string(l))
			list_part = append(list_part, i)
		}
		list = append(list, list_part)

	}
	list = list[0:(len(list) - 1)]
	return list
}
func abs(i int) int {
	if i > 0 {
		return i * 1
	} else {
		return i * -1
	}
}
func sum(is []int) int {
	sum := 0
	for _, l := range is {
		sum = sum + l
	}
	return sum
}
func asum(is []int) int {
	sum := 0
	for _, l := range is {
		sum = sum + abs(l)
	}

	return sum
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
	fmt.Println(report)
	answer := report_test(report)
	return answer
}
func report_test(iss [][]int) int {
	var success int
	for _, m := range iss {
		for k, l := range m {
			if (abs(l) <= 3) && (abs(l) >= 1) {

			} else {
				break
			}

			fmt.Println(k, m, l, abs(l))
			//its just counting every line right now
			//add a if k = len(m) or something
			success = success + 1
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
		o := bytes.Join(bytes.Split(l, []byte(" ")), []byte(""))
		for _, l := range o {
			i, _ := strconv.Atoi(string(l))
			list_part = append(list_part, i)
		}
		list = append(list, list_part)

	}
	list = list[0:(len(list) - 1)]
	return list
}
func abs(i int) int {
	if i >= 0 {
		return i * 1
	} else {
		return i * -1
	}
}

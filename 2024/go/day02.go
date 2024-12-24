package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Part 1: ", part1(data()))
	fmt.Println("Part 2: ", part2(data()), "649, 650 is wrong")
}
func part1(iss [][]int) int {
	var answer int
	iss = report(iss)

	for _, m := range iss {
		for k, l := range m {
			if (abs(sum(m)) != asum(m)) || (abs(l) > 3) || (l == 0) {
				break
			}
			if k+1 == len(m) {
				// fmt.Println(o, m, sum(m), asum(m), abs(sum(m)), l)
				answer = answer + 1
			}
		}
	}
	return answer
}
func part2(iss [][]int) int {
	var answer int
	var olen int
	iss = report(iss)

	for o, m := range iss {
		olen = len(m)
	loop:
		for k, l := range m {

			if (abs(sum(m)) != asum(m)) || (abs(l) > 3) || (l == 0) {
				if olen == len(m) {
					fmt.Println("start", o)
					fmt.Println("abs(sum()),asum()", abs(sum(m)), asum(m))
					fmt.Println("m:", m, "olen/len(m):", olen, len(m))
					fmt.Println("k:", k, "l:", l)
					fmt.Println("")
					m = slices.Delete(m, k, k+1)
					m = slices.Clip(m)

					continue loop
				} else if (abs(sum(m)) != asum(m)) || (abs(l) > 3) || (l == 0) {
					fmt.Println("end", o)
					fmt.Println("abs(sum()),asum()", abs(sum(m)), asum(m))
					fmt.Println("m:", m, "olen/len(m):", olen, len(m))
					fmt.Println("k:", k, "l:", l)
					fmt.Println("")
					break
				} else {
					fmt.Println(o, "point!")
					answer = answer + 1
					break
				}

			}
			if k+1 == len(m) {
				fmt.Println(o, "point!!")
				answer = answer + 1
			}
		}
	}
	return answer
}
func report(iss [][]int) [][]int {
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

	return report
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

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", part1(data()))

}
func part1(s string) int {
	var reg = regexp.MustCompile(`mul\(\d+\,\d+\)`)
	var num = regexp.MustCompile(`\d+\,\d+`)
	var com = regexp.MustCompile(`,`)

	answer := 0

	r := reg.FindAllString(s, -1)
	for _, l := range r {
		n := num.FindString(l)
		c := com.Split(n, -1)

		an1, _ := strconv.Atoi(c[0])
		an2, _ := strconv.Atoi(c[1])

		both := an1 * an2

		answer = both + answer
	}

	return answer
}
func data() string {
	ok, _ := os.ReadFile("data/day03.txt")
	data := string(ok)
	return data
}

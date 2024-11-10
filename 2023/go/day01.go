package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var answer_list []int

func main() {
	data, err := os.ReadFile("data/day01.txt")
	line := bytes.Split(data, []byte("\n"))

	if err != nil {
		panic(err)
	}

	for _, ele := range line {
		var answer_part []byte
		for _, ele_char := range ele {
			check := strings.ContainsAny(string(ele_char), "123456789")
			if check == true {
				answer_part = append(answer_part, ele_char)
			}
		}
		if len(answer_part) == 0 {

		} else {
			f := string(answer_part[0])
			l := string(answer_part[len(answer_part)-1])
			x, err := strconv.Atoi(strings.Join([]string{f, l}, ""))
			answer_list = append(answer_list, x)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(sumlist(answer_list))
}

func sumlist(x []int) int {
	sum := 0
	for _, ele := range x {
		sum += ele
	}
	return sum
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var answer_list []byte

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
		f := answer_part[0]
		l := answer_part[len(answer_part)-1]
		answer_list = append(answer_list, f, l)
	}
	fmt.Println(answer_list)
}

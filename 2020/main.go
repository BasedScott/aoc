package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data/day02.txt")
	list := bytes.Split(data, []byte("\n"))
	convert(list)
}

// func part1(bss [][]byte) int {
// 	return 0
// }

func convert(bss [][]byte) string {
	// var CONTAIN string
	var MIN string
	// var MAX int
	var PW string
	// var DEBUG string
	for _, bs := range bss {
		s := string(bs)
		_, PW, _ := strings.Cut(s, ": ")
		MIN, _, _ := strings.Cut(s, "-")
		fmt.Println("inside statement", "PW", PW, "MIN", MIN)
	}
	fmt.Println("outside statement", PW, MIN)
	return "ok"
}

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
	d := intslice
loop:
	for {
		for k, i := range d {
			if i+d[0] == 2020 {
				fmt.Println("Day01 :", i*d[0], i, d[0])
				break
			} else if k == len(d)-1 {
				d = slices.Delete(d, 0, 1)
				continue loop

			}
		}
		break
	}

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

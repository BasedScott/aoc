package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	var line = strings.Split(string(data), "\n")
	// var answer_part = []int{}

	for _, ele := range line {
		for _, ele_char := range ele {
			if err != nil {
				fmt.Println("Nope?", err)
				panic(err)
			}
			if int(ele_char) >= int(rune('9')) {
				fmt.Println("What is here?", string(ele_char), ele_char, int(rune('9')), len(ele))
			} else {
				fmt.Print("WAHT THE FAK", "\n")
			}
		}
		// make([]int, ele)
	}
}

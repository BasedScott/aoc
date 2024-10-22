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
	line := strings.Split(string(data), "\n")
	//var := 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 string
	//fmt.Println(len(line))
	for line := len(line)

}

package aoc

import (
	"fmt"
	"log"
	"os"
)

var m8 string

func main() {
	data, err := os.ReadFile("2023/data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func Get(m8 string) {
	os.Getenv(m8)
}

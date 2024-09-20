package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("2015/data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

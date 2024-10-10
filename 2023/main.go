package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

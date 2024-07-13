package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("/home/lenin/projects/aoc/2015/data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}

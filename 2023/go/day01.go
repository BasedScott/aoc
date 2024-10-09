package year2023

import (
	"fmt"
	"log"
	"os"
)

func main() {
	meme, err := os.ReadFile("/2023/data/day01.txt")
	if err != nil {
		log.Fatal(meme)
	}
	fmt.Println(meme)
}

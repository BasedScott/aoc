package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("data/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	var line = strings.Split(string(data), "\n")

	for _, cord := range line {
		re := regexp.MustCompile("[a-zA-Z]")
		fmt.Println(re.ReplaceAllString(cord, ""))

		//USE strings.NewReplacer for part 2
		// for _,each_cord_char := range each_cord {
		// 	fmt.Println(each_cord_char)
		// }
	}
}

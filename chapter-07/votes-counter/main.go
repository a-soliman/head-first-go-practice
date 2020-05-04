package main

import (
	"fmt"
	"log"

	"github.com/a-soliman/head_first_go/chapter-05/datafile"
)

func main() {
	var votes = make(map[string]int)

	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		votes[line]++
	}

	for name, value := range votes {
		fmt.Printf("%s : %d\n", name, value)
	}
}

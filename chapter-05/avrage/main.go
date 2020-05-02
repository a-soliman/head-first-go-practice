package main

import (
	"fmt"
	"log"

	"github.com/a-soliman/head_first_go/chapter-05/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	avrage := sum / float64(len(numbers))
	fmt.Printf("The avrage is = %.2f \n", avrage)
}

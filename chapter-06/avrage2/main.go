package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var numbers []float64

	for _, val := range args {
		number, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// numbers, err := datafile.GetFloats("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	avrage := sum / float64(len(numbers))
	fmt.Printf("The avrage is = %.2f \n", avrage)
}

package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	avrage := sum / float64(len(numbers))
	fmt.Printf("The avrage is = %.2f \n", avrage)
}

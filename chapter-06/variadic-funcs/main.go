package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	max := maximum(1, 2, 5, 4, 3)
	fmt.Print(max)
}

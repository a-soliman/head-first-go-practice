package main

import (
	"errors"
	"fmt"
)

func getArea(width float64, height float64) (area float64, err error) {
	if width < 0 {
		err = errors.New("Width can't be a negative value.")
		return
	}
	if height < 0 {
		err = errors.New("Height can't be a negative value.")
		return
	}
	area = width * height
	return
}

func calculatePaint(area float64) float64 {
	return area / 10.0
}

func print(paintAmount float64) {
	fmt.Printf("%.2f liters needed\n", paintAmount)
}

func paintNeeded(width float64, height float64) {
	area, err := getArea(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	paint := calculatePaint(area)
	print(paint)
}
func main() {
	paintNeeded(4.2, -3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, -3.3)
}

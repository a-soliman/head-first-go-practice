package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file
func GetFloats(filename string) ([3]float64, error) {
	var numbers [3]float64

	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

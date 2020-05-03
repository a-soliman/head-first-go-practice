package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file
func GetFloats(filename string) ([]float64, error) {
	var numbers []float64

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
		i++
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

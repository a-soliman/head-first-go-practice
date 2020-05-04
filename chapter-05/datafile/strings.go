package datafile

import (
	"bufio"
	"os"
)

// GetStrings reads a string from each line of a file
func GetStrings(filename string) ([]string, error) {
	var lines []string
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// initalite the scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// close the file
	err = file.Close()
	if err != nil {
		return nil, err
	}

	// check for scanner errors
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	// return the lines
	return lines, nil
}

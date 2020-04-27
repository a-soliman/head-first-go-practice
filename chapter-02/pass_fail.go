package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("a grade of ", grade, ", has a status of ", status)
}

// ask the user to provide an input
// read the input usinbg bufio
// convert the input into float 64
// check if failing or passing grade

package main

import (
	"fmt"
	"log"

	"github.com/a-soliman/keyboard"
)

func main() {
	fmt.Print("Enter a grade : ")

	grade, err := keyboard.GetFloat()
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

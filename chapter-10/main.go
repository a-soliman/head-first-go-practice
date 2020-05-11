package main

import (
	"log"

	"github.com/a-soliman/head_first_go/calendar"
)

func main() {
	var myBirthDay calendar.Event
	var err error
	err = myBirthDay.SetTitle("Ahmed's Birthday")
	err = myBirthDay.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = myBirthDay.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = myBirthDay.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}
	myBirthDay.Print()
}

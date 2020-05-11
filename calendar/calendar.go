package calendar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Event struct
type Event struct {
	title string
	Date
}

// Title getter
func (e *Event) Title() string {
	return e.title
}

// SetTitle setter
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid title")
	}
	e.title = title
	return nil
}

// Print prints an event
func (e *Event) Print() {
	fmt.Printf("Title: %s\n", e.Title())
	e.Date.Print()
}

// Date struct
type Date struct {
	year  int
	month int
	day   int
}

// Year getter
func (d *Date) Year() int {
	return d.year
}

// Month getter
func (d *Date) Month() int {
	return d.month
}

// Day getter
func (d *Date) Day() int {
	return d.day
}

// SetYear setter
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

// SetMonth setter
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

// SetDay setter
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}

// Print prints the date instence
func (d *Date) Print() {
	fmt.Printf("%d - %d -%d.\n", d.year, d.month, d.day)
}

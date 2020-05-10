package magazine

import "fmt"

// Subscriber struct defines a blue print for the subscriber
type Subscriber struct {
	name   string
	rate   float64
	active bool
	Address
}

// Employee struct defines a blue print for the employee
type Employee struct {
	Name   string
	Salary float64
	Address
}

// Address struct type
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// DefaultSubscriber returns a pointer to a newly created subscriber
func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.name = name
	s.rate = 5.99
	s.active = true

	return &s
}

// PrintInfo prints the fields of a subscriber
func (s *Subscriber) PrintInfo() {
	fmt.Println("Name: ", s.name)
	fmt.Println("Rate: ", s.rate)
	fmt.Println("active? ", s.active)
	fmt.Printf("Address:\nStreet: %s,\n%s, %s\n%s", s.Street, s.City, s.State, s.PostalCode)
}

// ApplyDiscount applys discount on a given subscriber
func (s *Subscriber) ApplyDiscount() {
	s.rate = 4.99
}

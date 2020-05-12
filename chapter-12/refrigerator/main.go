package main

import (
	"errors"
	"fmt"
	"log"
)

type refrigerator struct {
	items    []string
	isOpened bool
}

func (r *refrigerator) open() {
	fmt.Println("Opening.")
	r.isOpened = true
}

func (r *refrigerator) close() {
	fmt.Println("Closing...")
	r.isOpened = false
}

func (r *refrigerator) addItems(items []string) {
	for _, item := range items {
		r.addItem(item)
	}
}

func (r *refrigerator) addItem(item string) {
	r.items = append(r.items, item)
}

func (r *refrigerator) find(item string) (string, error) {
	if r.isOpened != true {
		r.open()
		defer r.close()
	}
	for _, rItem := range r.items {
		if item == rItem {
			return rItem, nil
		}
	}
	return "", errors.New("item not found")
}

func main() {
	fridge := refrigerator{}
	fridge.addItems([]string{"milk", "meat", "juice"})
	item, err := fridge.find("milk")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(item)

}

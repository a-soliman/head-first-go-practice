package main

import (
	"github.com/a-soliman/head_first_go/magazine"
)

func main() {
	subscriber1 := magazine.DefaultSubscriber("Ahmed Soliman")
	subscriber1.PrintInfo()
	subscriber1.ApplyDiscount()
	subscriber1.PrintInfo()
}

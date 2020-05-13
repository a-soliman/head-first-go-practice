package main

import (
	"fmt"

	"github.com/a-soliman/head_first_go/prose"
)

func main() {
	test := prose.JoinWithCommas([]string{"someone", "was", "here"})
	fmt.Println(test)
}

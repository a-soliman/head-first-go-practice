package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readDirectory(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, file := range files {
		path := dirname + "/" + file.Name()
		if file.IsDir() {
			fmt.Println("Directory: ", path)
			readDirectory(path)
		} else {
			fmt.Println("File: ", path)
		}
	}
	return nil
}

func main() {
	err := readDirectory("my_directory")
	if err != nil {
		log.Fatal(err)
	}
}

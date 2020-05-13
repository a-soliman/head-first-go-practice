package main

import (
	"log"
	"net/http"
)

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello world")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

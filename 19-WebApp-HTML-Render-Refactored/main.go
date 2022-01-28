package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8081"

func main() {
	fmt.Printf("This is a sample line")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting Application on Port %v", portNumber)

	http.ListenAndServe(portNumber, nil)
}

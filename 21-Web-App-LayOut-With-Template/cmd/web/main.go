package main

import (
	"fmt"
	"net/http"

	"github.com/VargheseVibin/GoWebPackage01/pkg/handlers"
)

const portNumber = ":8081"

func main() {
	fmt.Printf("This is a sample line")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting Application on Port %v", portNumber)

	http.ListenAndServe(portNumber, nil)
}

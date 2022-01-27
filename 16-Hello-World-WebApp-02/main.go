package main

import (
	"fmt"
	"io"
	"net/http"
)

const portNumber = ":8081"

// Home if the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the HOME Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addNumbers(6, 7)
	io.WriteString(w, fmt.Sprintf("This is about me page telling me 7 + 6 is %d", sum))
}

func addNumbers(a int, b int) int {
	return a + b
}

func main() {
	fmt.Printf("This is a sample line")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting Application on Port %v", portNumber)

	http.ListenAndServe(portNumber, nil)
}

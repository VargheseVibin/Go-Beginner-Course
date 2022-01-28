package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	var a, b float32
	a, b = 100.0, 0.0
	f, err := divideValues(a, b)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
	} else {
		// fmt.Fprintf(w, fmt.Sprintf("Division: %f divided by %f is %f", 100.0, 10.0, f))
		io.WriteString(w, fmt.Sprintf("Division: %f divided by %f is %f", a, b, f))
	}
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot Divide by zero. check divisor")
		return 0, err
	}
	result := x / y
	return result, nil
}
func main() {
	fmt.Printf("This is a sample line")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting Application on Port %v", portNumber)

	http.ListenAndServe(portNumber, nil)
}

package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(120.0, 0)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("result of division is ", result)
	}

}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot Divide by 0!")
	}

	result = x / y
	return result, nil
}

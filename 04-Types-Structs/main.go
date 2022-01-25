package main

import (
	"log"
	"time"
)

// variable Names with first letter as upper case can be referred from different packages
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Stephen",
		LastName:    "Colbert",
		PhoneNumber: "1-123-456-7890",
	}
	log.Println(user.FirstName, user.LastName, "has a birthdate of:", user.BirthDate)
}

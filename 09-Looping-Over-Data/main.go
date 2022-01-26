package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "bird", "cat", "horse"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"
	animalsMap["fish"] = "Nemo"
	animalsMap["lion"] = "Simba"

	for animalType, animalName := range animalsMap {
		log.Println(animalType, animalName)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Vibin", "Varghese", "vv@g.com", 35})
	users = append(users, User{"Febin", "Varghese", "fv@g.com", 45})
	users = append(users, User{"Anu", "Baby", "ab@g.com", 32})
	users = append(users, User{"Anu", "Jacob", "aj@g.com", 40})

	for _, usr := range users {
		log.Println(usr.FirstName, usr.LastName, usr.Email, usr.Age)
	}
}

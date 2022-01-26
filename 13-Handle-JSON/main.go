package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Player struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Country       string `json:"country"`
	Plays         string `json:"plays"`
	GrandSlamsWon int    `json:"grand_slams"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Roger",
			"last_name": "Federer",
			"country": "SUI",
			"plays": "Right-Handed",
			"grand_slams": 20
		},
		{
			"first_name": "Rafael",
			"last_name": "Nadal",
			"country": "ESP",
			"plays": "Left-Handed",
			"grand_slams": 20
		},
		{
			"first_name": "Pete",
			"last_name": "Sampras",
			"country": "USA",
			"plays": "Right-Handed",
			"grand_slams": 14
		}
	]
	`

	var unmarshalled []Player

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error Unmarshalling JSON", err)
	}
	log.Printf("unmarshalled %v", unmarshalled)

	p1 := Player{"Rod", "Laver", "AUS", "Left-Handed", 11}
	p2 := Player{"Stefan", "Edberg", "SWE", "Right-Handed", 11}

	var myPlayers []Player
	myPlayers = append(myPlayers, p1)
	myPlayers = append(myPlayers, p2)

	newPlayersJson, err := json.MarshalIndent(myPlayers, "", "	")
	if err != nil {
		log.Println("error marshalling ", err)
	}

	fmt.Println(string(newPlayersJson))

}

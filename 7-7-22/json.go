package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func Json() {
	myJson := `
	[
		{
			"first_name": "Clark"
		},
		{}
	]
`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Default().Println("Error unmarshalling json")
	}

	log.Printf("unmarshalled: %v", unmarshalled)
}

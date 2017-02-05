package main

import (
	"encoding/json"
	"net/http"

	"github.com/WWGLondon/ExampleAPI/animals"
)

var pets []animals.Pet

func main() {
	pets = []animals.Pet{
		&animals.Kitten{
			Name: "Mr Tiggles",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
		&animals.Kitten{
			Name: "Top Cat",
			Hobbies: []string{
				"Taunting Officer Dibble",
			},
		},
		&animals.Dog{
			Name: "Fido",
			Hobbies: []string{
				"Barking",
			},
		},
	}

	http.HandleFunc("/list", ListKittens)

	// Homework B.4
	// TODO: Register a new "/like" route and associate it with `LikedPet` handler function

	http.ListenAndServe(":9000", http.DefaultServeMux)
}

func ListKittens(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(pets)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}

// Homework B.5
// Modify this handler function to be able to handle `localhost:9000/like?name=Fido`
func LikedPet(rw http.ResponseWriter, r *http.Request) {
	// Homework B.5.1
	// TODO:
	// assign a new variable "name" and it should be of type string
	// parse query string with "name" key and assign the value to name variable

	// Homework B.5.2
	// TODO:
	// iterate through pets
	// increment Like value of a pet only if the pet's name equal to variable name
}

// go run main.go

package main

import (
	"encoding/json"
	"net/http"

	"github.com/nicholasjackson/wwg/animals"
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

// go run main.go

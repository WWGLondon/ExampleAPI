package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Pet struct {
	Name    string
	Hobbies []string
}

func main() {
	resp, err := http.Get("http://localhost:9000/list")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pets := []Pet{}
	err = json.Unmarshal(data, &pets)

	//loop through pets and print each one on a line
}

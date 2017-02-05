package animals

import (
	"reflect"
	"testing"
)

// go test ./...
func TestSetHobbiesSetsTheHobbies(t *testing.T) {
	//setup
	hobbies := []string{"Barking"}
	dog := Dog{}

	//do work
	dog.SetHobbies(hobbies)

	//assert
	if !reflect.DeepEqual(dog.GetHobbies(), hobbies) {
		t.Fail()
	}
}

func TestGetHobbiesGetsTheHobbies(t *testing.T) {
	// setup
	hobbies := []string{"Barking"}
	dog := Dog{Hobbies: hobbies}

	// assert
	if dog.GetHobbies()[0] != hobbies[0] {
		t.Fail()
	}
}

func TestSetNameSetsTheName(t *testing.T) {
	// setup
	dog := Dog{}

	// do work
	dog.SetName("Spots")

	// assert
	if dog.Name != "Spots" {
		t.Fail()
	}
}

func TestGetNameGetsTheName(t *testing.T) {
	// setup
	dog := Dog{Name: "Spots"}

	// assert
	if dog.GetName() != "Spots" {
		t.Fail()
	}
}

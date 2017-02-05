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

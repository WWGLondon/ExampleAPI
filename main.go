package main

import (
	"fmt"

	"github.com/nicholasjackson/wwg/animals"
)

func main() {
	kitty := animals.Kitten{}
	kitty.SetName("Mr Tiggles")

	fmt.Println(kitty.GetName())

	doggy := animals.Dog{}
	fmt.Println(doggy.Bark())
}

// go run main.go

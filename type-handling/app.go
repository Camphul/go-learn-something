package main

import "log"

type Animal struct {
	Name         string
	ProduceSound func()
}

type Dog = Animal
type Cat = Animal

var goldenRetriever = Dog{"Jeff", func() {
	log.Println("Hello I am a golden retriever")
}}
var poodle = Dog{"Deez", func() {
	log.Println("Hello I am a poodle")
}}
var kot = Cat{"Concrete", func() {
	log.Println("I a kot lol")
}}

func main() {
	log.Println("Running type-handling program..")
	pets := [3]Animal{goldenRetriever, poodle, kot}
	log.Printf("Amount of pets: %s\n", len(pets))
	for _, pet := range pets {
		log.Printf("Pet %s says:\n", pet.Name)
		pet.ProduceSound()
	}
}

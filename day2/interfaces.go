// Interfaces
// Sets of Methods (A type implements an interface by implementing its interfaces, Theyt are central to go type systems and polymorphism)
package main

import "fmt"

// Pet interface defines behavior that pets must have
type Pet interface {
	Speak() string // Any type that has a Speak method is an Pet
}

// Dog type represents a dog
type Dog struct{}

// Cat type represents a cat
type Cat struct{}

// Speak for Dog returns a bhoww sound
func (d Dog) Speak() string {
	return "Bhowww!!!"
}

// Speak for Cat returns a meowing sound
func (c Cat) Speak() string {
	return "Meow"
}

func main() {

	// Create a slice of Pet interface, holding different types of animals
	animals := []Pet{Dog{}, Cat{}}

	// Loop through animals and call their Speak method
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}

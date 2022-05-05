package main

import (
	"fmt"
	"time"

	animal "github.com/braejan/learning-golang-poo/src/domain/animal/interfaces"
	"github.com/braejan/learning-golang-poo/src/domain/cat/entities"
	cat "github.com/braejan/learning-golang-poo/src/domain/cat/interfaces"
	"github.com/braejan/learning-golang-poo/src/domain/cat/usecases"
)

func main() {
	fmt.Println("Wellcome to POO in Golang.")
	catHoney := &entities.Cat{
		Name:       "Honey",
		BirthDay:   time.Now(),
		Sterilized: true,
	}
	animal := usecases.NewCatImpl(catHoney)

	RunAllAnimalMethods(animal)

	DrinkYourWater(animal, "milk")
}

func RunAllAnimalMethods(animal animal.AnimalI) {
	fmt.Println("move: ", animal.Move())
	fmt.Println("feed:", animal.Feed())
}

func DrinkYourWater(cat cat.CatI, drink string) {
	cat.Drink(drink)
}

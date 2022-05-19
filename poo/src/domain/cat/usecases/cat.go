package usecases

import (
	"fmt"
	"math/rand"

	"github.com/braejan/learning-golang-poo/src/domain/cat/entities"
)

type CatUsecases struct {
	*entities.Cat
}

func NewCatUsecases(cat *entities.Cat) (usecase *CatUsecases) {
	usecase = &CatUsecases{
		Cat: cat,
	}
	return
}

func (usecase *CatUsecases) Feed() (result bool) {
	food := "pepitas"
	drink := "water"
	result = usecase.Drink(drink) && usecase.Eat(food)
	return
}
func (usecase *CatUsecases) Breathe() (result bool) {
	fmt.Printf("%s breathe ... and then exhale.\n", usecase.Name)
	result = true
	return
}
func (usecase *CatUsecases) Move() (result bool) {
	lazy := rand.Intn(2) + 1
	if lazy == 1 {
		return usecase.Walk()
	}
	return usecase.Run()
}
func (usecase *CatUsecases) Reproduce() (result bool) {
	result = !usecase.Sterilized
	return
}
func (usecase *CatUsecases) Eat(food string) (result bool) {
	fmt.Printf("%s: Yumy %s Yumy\n", usecase.Name, food)
	result = true
	return
}
func (usecase *CatUsecases) Drink(drink string) (result bool) {
	fmt.Printf("%s: Tasty %s\n", usecase.Name, drink)
	result = true
	return
}
func (usecase *CatUsecases) Walk() (result bool) {
	fmt.Printf("%s: I'm walking 'cause I feel lazy today.\n", usecase.Name)
	result = true
	return
}
func (usecase *CatUsecases) Run() (result bool) {
	fmt.Printf("%s: I'm runnin 'cause I feel crazy today.\n", usecase.Name)
	result = true
	return
}
func (usecase *CatUsecases) Purr() {
	for i := 0; i < int(usecase.VoiceLevel); i++ {
		fmt.Printf("%s: prrrrrrprrrrrrprrrrr\n", usecase.Name)
	}
	fmt.Println(" -> :)")
}
func (usecase *CatUsecases) Meow() {
	if usecase.VoiceLevel > 5 {
		fmt.Printf("%s terribol: Aaaaauuuuuuuuuuuuuuuuuuuuuuu\n", usecase.Name)
	} else {
		fmt.Printf("%s: Meowww Meowwwww", usecase.Name)
	}
}

package cat

import animal "github.com/braejan/learning-golang-poo/src/domain/animal/interfaces"

type CatI interface {
	animal.AnimalI
	Eat(food string) (result bool)
	Drink(food string) (result bool)
	Walk() (result bool)
	Run() (result bool)
	Purr(love int) (result bool)
	Meow() (result bool)
}

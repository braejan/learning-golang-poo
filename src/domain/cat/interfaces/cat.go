package interfaces

import (
	"github.com/braejan/learning-golang-poo/src/domain/animal/interfaces"
)

type LivingCat interface {
	interfaces.Animal //No hay herencia, sólo composición
	Eat(food string) (result bool)
	Drink(drink string) (result bool)
	Walk() (result bool)
	Run() (result bool)
	Purr()
	Meow()
}

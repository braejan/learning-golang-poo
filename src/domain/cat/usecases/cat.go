package usecases

import (
	"fmt"

	"github.com/braejan/learning-golang-poo/src/domain/cat/entities"
)

type CatImpl struct {
	*entities.Cat
}

func NewCatImpl(cat *entities.Cat) (catImpl *CatImpl) {
	catImpl = &CatImpl{
		Cat: cat,
	}
	return
}

func (c *CatImpl) Breathe() (result bool) {
	result = true
	return
}
func (c *CatImpl) Feed() (result bool) {
	result = c.Eat("pepitas") && c.Drink("agüita")
	return
}
func (c *CatImpl) Move() (result bool) {
	result = c.Run() || c.Walk()
	return
}
func (c *CatImpl) Reproduce() (result bool) {
	result = !c.Sterilized
	return
}

func (c *CatImpl) Eat(food string) (result bool) {
	fmt.Printf("%s: que rico/a %s\n", c.Name, food)
	result = true
	return
}
func (c *CatImpl) Drink(food string) (result bool) {
	fmt.Printf("%s: yumy yumy %s\n", c.Name, food)
	result = true
	return
}
func (c *CatImpl) Walk() (result bool) {
	fmt.Printf("%s: está caminando juicioso/a.\n", c.Name)
	result = true
	return
}
func (c *CatImpl) Run() (result bool) {
	fmt.Printf("%s: está corriendo como loco.\n", c.Name)
	result = true
	return
}
func (c *CatImpl) Purr(love int) (result bool) {
	if love > 50 {
		fmt.Printf("%s: prrrr prrrr prrr prrrr prrrr\n", c.Name)
		result = true
	}
	return
}
func (c *CatImpl) Meow() (result bool) {
	fmt.Printf("%s: Meow Meow Meow Meow Meow Meow Meow Meow Meow!!!!!\n", c.Name)
	result = true
	return
}

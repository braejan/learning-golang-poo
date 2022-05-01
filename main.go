package main

import (
	"fmt"
	"time"

	"github.com/braejan/learning-golang-poo/src/domain/cat/entities"
	"github.com/braejan/learning-golang-poo/src/domain/cat/usecases"
)

func main() {
	fmt.Println("Wellcome to learn POO in Golang.")
	/*
	   Abstracción
	   Encapsulamiento
	   Polimorfismo
	   Herencia
	   Modularidad
	   Principio de ocultación
	   Recolección de basura
	*/
	bruce := entities.NewCat("Bruce", new(time.Time), 8, "yellow")
	kitty := entities.NewCat("Kitty", new(time.Time), 4, "white")
	honey := entities.NewCat("Honey", new(time.Time), 5, "gray")
	arthur := entities.NewCat("Arthur", new(time.Time), 7, "gray")

	bruceUsecase := usecases.NewCatUsecases(bruce)
	kittyUsecase := usecases.NewCatUsecases(kitty)
	honeyUsecase := usecases.NewCatUsecases(honey)
	arthurUsecase := usecases.NewCatUsecases(arthur)

	bruceUsecase.Purr()
	kittyUsecase.Move()
	honeyUsecase.Feed()
	arthurUsecase.Run()
}

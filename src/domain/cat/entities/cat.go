package entities

import (
	"time"

	"github.com/braejan/learning-golang-poo/src/domain/cat/interfaces"
)

type Cat struct {
	interfaces.LivingCat
	Name           string
	BirthDate      *time.Time
	VoiceLevel     int8
	PrincipalColor string
	Sterilized     bool
}

func NewCat(name string, birthDate *time.Time, voiceLevel int8, color string) (cat *Cat) {
	cat = &Cat{
		Name:           name,
		BirthDate:      birthDate,
		VoiceLevel:     voiceLevel,
		PrincipalColor: color,
	}
	return
}

package entities

import "time"

type Cat struct {
	Name       string
	BirthDay   time.Time
	Color      string
	Sterilized bool
}

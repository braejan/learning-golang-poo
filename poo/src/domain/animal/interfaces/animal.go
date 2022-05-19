package interfaces

type Animal interface {
	Feed() (result bool)
	Breathe() (result bool)
	Move() (result bool)
	Reproduce() (result bool)
}

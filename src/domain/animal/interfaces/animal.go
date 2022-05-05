package animal

type AnimalI interface {
	Breathe() (result bool)
	Feed() (result bool)
	Move() (result bool)
	Reproduce() (result bool)
}

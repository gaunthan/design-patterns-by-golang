package factorymethod

type ButtonFactory interface {
	New() Button
}

type CircularButtonFactory struct{}

func (f CircularButtonFactory) New() Button {
	return new(CircularButton)
}

type SquareButtonFactory struct{}

func (f SquareButtonFactory) New() Button {
	return new(SquareButton)
}

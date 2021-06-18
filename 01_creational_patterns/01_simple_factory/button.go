package simplefactory

type Button interface {
	Press() string
}

type CircularButton struct{}

func (b CircularButton) Press() string {
	return "pressing circular button"
}

type SquareButton struct{}

func (b SquareButton) Press() string {
	return "pressing square button"
}

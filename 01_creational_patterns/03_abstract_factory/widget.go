package abstractfactory

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

type ScrollBar interface {
	Scroll() string
}

type CircularScrollBar struct{}

func (b CircularScrollBar) Scroll() string {
	return "scrolling circular scroll bar"
}

type SquareScrollBar struct{}

func (b SquareScrollBar) Scroll() string {
	return "scrolling square scroll bar"
}

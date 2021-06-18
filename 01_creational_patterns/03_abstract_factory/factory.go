package abstractfactory

type WidgetFactory interface {
	NewButton() Button
	NewScrollBar() ScrollBar
}

type CircularWidgetFactory struct{}

func (f CircularWidgetFactory) NewButton() Button {
	return new(CircularButton)
}

func (f CircularWidgetFactory) NewScrollBar() ScrollBar {
	return new(CircularScrollBar)
}

type SquareWidgetFactory struct{}

func (f SquareWidgetFactory) NewButton() Button {
	return new(SquareButton)
}

func (f SquareWidgetFactory) NewScrollBar() ScrollBar {
	return new(SquareScrollBar)
}

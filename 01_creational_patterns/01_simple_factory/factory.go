package simplefactory

type ButtonFactory struct{}

func (f ButtonFactory) New(buttonType string) Button {
	switch buttonType {
	case "circular":
		return new(CircularButton)
	case "square":
		return new(SquareButton)
	default:
		return nil
	}
}

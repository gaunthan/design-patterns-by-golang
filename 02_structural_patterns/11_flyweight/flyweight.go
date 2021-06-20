package flyweight

import "fmt"

type ImageFlyweight interface {
	Data() string
}

type ConcreteImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) ImageFlyweight {
	return &ConcreteImageFlyweight{
		data: openAndReadAll(filename),
	}
}

func (p *ConcreteImageFlyweight) Data() string {
	return p.data
}

func openAndReadAll(filename string) string {
	return fmt.Sprintf("data of %s", filename)
}

package bridge

import "fmt"

type Sender interface {
	Send(str string) string
}

type DogSender struct{}

func (*DogSender) Send(str string) string {
	return fmt.Sprintf("[Dog] %s", str)
}

func NewDogSender() *DogSender {
	return &DogSender{}
}

type CatSender struct{}

func (*CatSender) Send(str string) string {
	return fmt.Sprintf("[Cat] %s", str)
}

func NewCatSender() *CatSender {
	return &CatSender{}
}

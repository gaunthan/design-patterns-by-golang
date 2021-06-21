package observer

import "fmt"

type Observer interface {
	OnUpdate(msg string)
}

func NewObserver(name string) Observer {
	return &observerImpl{
		name: name,
	}
}

type observerImpl struct {
	name string
}

func (p *observerImpl) OnUpdate(msg string) {
	fmt.Printf("%s: received %q\n", p.name, msg)
}

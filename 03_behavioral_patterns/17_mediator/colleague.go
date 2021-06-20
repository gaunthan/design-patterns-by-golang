package mediator

import "fmt"

type Colleague interface {
	Name() string
	Speak(to, content string)
	OnListen(from, content string)
	SetMediator(mediator Mediator)
}

type colleagueImpl struct {
	name     string
	mediator Mediator
}

func NewColleague(name string) Colleague {
	return &colleagueImpl{
		name: name,
	}
}

func (p *colleagueImpl) Name() string {
	return p.name
}

func (p *colleagueImpl) Speak(to, content string) {
	p.mediator.Whisper(p.name, to, content)
}

func (p *colleagueImpl) OnListen(from, content string) {
	fmt.Printf("%s: I heard %q from %s\n", p.name, content, from)
}

func (p *colleagueImpl) SetMediator(mediator Mediator) {
	p.mediator = mediator
}

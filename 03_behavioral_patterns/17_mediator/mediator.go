package mediator

type Mediator interface {
	Whisper(from, to, content string)
	Register(colleague Colleague)
}

type mediatorImpl struct {
	contacts map[string]Colleague
}

func NewMediator() Mediator {
	return &mediatorImpl{
		contacts: make(map[string]Colleague),
	}
}

func (p *mediatorImpl) Whisper(from, to, content string) {
	// if the target not found, whisper to everyone
	if _, ok := p.contacts[to]; !ok {
		names := getNames(p.contacts)
		for _, name := range names {
			if name == from {
				continue
			}
			p.contacts[name].OnListen(from, content)
		}
	} else { // else to specified target
		p.contacts[to].OnListen(from, content)
	}
}

func (p *mediatorImpl) Register(colleagure Colleague) {
	p.contacts[colleagure.Name()] = colleagure
	colleagure.SetMediator(p)
}

func getNames(contacts map[string]Colleague) []string {
	names := make([]string, 0, len(contacts))
	for name := range contacts {
		names = append(names, name)
	}
	return names
}

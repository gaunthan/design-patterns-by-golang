package observer

import "container/list"

type Subject interface {
	attach(observer Observer)
	detach(observer Observer)
	notify(msg string)
}

func NewSubject() Subject {
	return &subjectImpl{
		observers: list.New(),
		maps:      make(map[Observer]*list.Element),
	}
}

type subjectImpl struct {
	observers *list.List
	maps      map[Observer]*list.Element
}

func (p *subjectImpl) attach(observer Observer) {
	el := p.observers.PushBack(observer)
	p.maps[observer] = el
}

func (p *subjectImpl) detach(observer Observer) {
	if _, ok := p.maps[observer]; !ok {
		return
	}
	el := p.maps[observer]
	p.observers.Remove(el)
	delete(p.maps, observer)
}

func (p *subjectImpl) notify(msg string) {
	for e := p.observers.Front(); e != nil; e = e.Next() {
		e.Value.(Observer).OnUpdate(msg)
	}
}

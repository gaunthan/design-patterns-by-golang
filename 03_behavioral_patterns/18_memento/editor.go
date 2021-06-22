package memento

type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (p *Editor) Type(input string) {
	p.content += input
}

func (p *Editor) Content() string {
	return p.content
}

func (p *Editor) Save() Memento {
	memento := NewEditorMemento()
	memento.SetState(p.content)
	return memento
}

func (p *Editor) Restore(memento Memento) {
	p.content = memento.GetState().(string)
}

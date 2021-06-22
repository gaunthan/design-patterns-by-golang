package memento

type Memento interface {
	SetState(state interface{})
	GetState() interface{}
}

type EditorMemento struct {
	content string
}

func NewEditorMemento() *EditorMemento {
	return &EditorMemento{}
}

func (p *EditorMemento) SetState(state interface{}) {
	p.content = state.(string)
}

func (p *EditorMemento) GetState() interface{} {
	return p.content
}

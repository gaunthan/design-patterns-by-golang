package state

type Context struct {
	state State
}

func NewContext() *Context {
	return &Context{
		state: BeginState,
	}
}

func (p *Context) ToUpper(input []rune) {
	index := 0
	for p.state != EndState {
		p.state = p.state.handle(input, index)
		index++
	}
}

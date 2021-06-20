package command

import "fmt"

type Command interface {
	Execute()
}

var values = map[string]interface{}{}

type SetIntCommand struct {
	key   string
	value int
}

func NewSetIntCommand(key string, value int) *SetIntCommand {
	return &SetIntCommand{
		key:   key,
		value: value,
	}
}

func (p *SetIntCommand) Execute() {
	values[p.key] = p.value
}

type PrintValueCommand struct {
	key string
}

func NewPrintValueCommand(key string) *PrintValueCommand {
	return &PrintValueCommand{
		key: key,
	}
}

func (p *PrintValueCommand) Execute() {
	fmt.Println(values[p.key])
}

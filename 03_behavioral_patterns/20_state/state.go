package state

import (
	"unicode"
)

func init() {
	BeginState = &beginState{}
	EndState = &endState{}
}

type State interface {
	handle(input []rune, index int) State
}

var (
	BeginState State
	EndState   State
)

type beginState struct{}

func (*beginState) handle(input []rune, index int) State {
	if index >= len(input) {
		return EndState
	}
	input[index] = unicode.ToUpper(input[index])
	return BeginState
}

type endState struct{}

func (p *endState) handle(input []rune, index int) State {
	return EndState
}

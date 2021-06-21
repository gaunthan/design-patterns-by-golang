package state

import (
	"unicode"
)

func init() {
	BeginState = &beginState{}
	EndState = &endState{}
	ToUpperState = &toUpperState{}
	ToLowerState = &toLowerState{}
}

type State interface {
	handle(input []rune, index int) State
}

var (
	BeginState   State
	EndState     State
	ToUpperState State
	ToLowerState State
)

type beginState struct{}

func (*beginState) handle(input []rune, index int) State {
	if index >= len(input) {
		return EndState
	}
	return ToUpperState
}

type toUpperState struct{}

func (*toUpperState) handle(input []rune, index int) State {
	if index >= len(input) {
		return EndState
	}
	input[index] = unicode.ToUpper(input[index])
	return ToLowerState
}

type toLowerState struct{}

func (*toLowerState) handle(input []rune, index int) State {
	if index >= len(input) {
		return EndState
	}
	input[index] = unicode.ToLower(input[index])
	return ToUpperState
}

type endState struct{}

func (p *endState) handle(input []rune, index int) State {
	return EndState
}

package adapter

import "fmt"

type Screen struct {
	in DVIPort
}

func (s *Screen) SetInput(in DVIPort) {
	s.in = in
}

func (s *Screen) Play() {
	for msg := range s.in {
		fmt.Print(msg)
	}
}

package adapter

import "fmt"

type HDMI2DVIAdapter struct {
	out DVIPort
}

func (a *HDMI2DVIAdapter) SetInput(in HDMIPort) {
	a.out = make(DVIPort)
	go func() {
		defer close(a.out)

		for msg := range in {
			a.out <- fmt.Sprint(msg)
		}
	}()
}

func (a *HDMI2DVIAdapter) GetOutput() DVIPort {
	return a.out
}

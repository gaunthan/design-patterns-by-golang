package adapter

type VideoOutputter struct {
	out HDMIPort
}

func (o *VideoOutputter) Start() {
	o.out = make(HDMIPort)
	go func() {
		defer close(o.out)

		// outputs 0 1 2 ... 9
		for i := 0; i < 10; i++ {
			o.out <- i
		}
	}()
}

func (o *VideoOutputter) GetOutput() HDMIPort {
	return o.out
}

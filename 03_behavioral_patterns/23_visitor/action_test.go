package visitor

func Example_one() {
	car := NewAutomobile()
	moto := NewMotocycle()

	forward := NewMoveForwardAction()
	car.Accept(forward)
	moto.Accept(forward)

	drift := NewDriftAction()
	car.Accept(drift)
	moto.Accept(drift)

	// Output:
	// Automobile is moving forward fastly
	// Motocycle is moving forward slowly
	// Automobile is drifting stably
	// Motocycle is drifting carefully
}

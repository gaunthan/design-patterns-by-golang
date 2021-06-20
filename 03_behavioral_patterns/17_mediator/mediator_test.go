package mediator

func Example_one() {
	mediator := NewMediator()

	jack := NewColleague("jack")
	mediator.Register(jack)

	tom := NewColleague("tom")
	mediator.Register(tom)

	bob := NewColleague("bob")
	mediator.Register(bob)

	jack.Speak("tom", "hi")
	tom.Speak("bob", "hey")

	// Output:
	// tom: I heard "hi" from jack
	// bob: I heard "hey" from tom
}

func Example_two() {
	mediator := NewMediator()

	jack := NewColleague("jack")
	mediator.Register(jack)

	tom := NewColleague("tom")
	mediator.Register(tom)

	bob := NewColleague("bob")
	mediator.Register(bob)

	jack.Speak("", "hi")

	// Output:
	// tom: I heard "hi" from jack
	// bob: I heard "hi" from jack
}

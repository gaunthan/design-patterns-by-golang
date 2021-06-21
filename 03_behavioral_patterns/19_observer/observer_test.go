package observer

import "fmt"

func ExampleObserver() {
	weather := NewSubject()
	joe := NewObserver("joe")
	weather.attach(joe)
	tom := NewObserver("tom")
	weather.attach(tom)

	weather.notify("today is sunshiny")
	fmt.Println("---")

	weather.detach(joe)
	weather.notify("tomorrow is cloudy")
	fmt.Println("---")

	weather.detach(tom)
	weather.notify("oh my god")

	// Output:
	// joe: received "today is sunshiny"
	// tom: received "today is sunshiny"
	// ---
	// tom: received "tomorrow is cloudy"
	// ---
	//
}

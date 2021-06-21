package state

import "fmt"

func ExampleState() {
	context := NewContext()

	input := []rune("HELLO, world")
	context.Transform(input)

	fmt.Println(string(input))

	// Output:
	// HElLo, WoRlD
}

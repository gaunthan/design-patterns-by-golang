package state

import "fmt"

func ExampleState() {
	context := NewContext()

	input := []rune("Hello, world!")
	context.ToUpper(input)

	fmt.Println(string(input))

	// Output:
	// HELLO, WORLD!
}

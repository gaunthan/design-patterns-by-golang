package memento

import "fmt"

func Example_one() {
	editor := NewEditor()
	editor.Type("Hello")
	editor.Type(", world")

	memento := editor.Save()
	editor.Type("sjifjsifni")
	editor.Restore(memento)

	fmt.Println(editor.Content())

	// Output:
	// Hello, world
}

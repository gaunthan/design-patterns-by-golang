package iterator

import "fmt"

func Example_one() {
	vector := NewVector()
	vector.Push(1)
	vector.Push(2)
	vector.Push(3)

	for iter := vector.Begin(); iter != vector.End(); iter = iter.Next() {
		value := iter.Data().(int)
		fmt.Println(value)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_two() {
	vector := NewVector()
	vector.Push(1)
	vector.Push(2)
	vector.Push(3)
	vector.Push(4)
	vector.Pop()

	for iter := vector.RightBegin(); iter != vector.RightEnd(); iter = iter.Prev() {
		value := iter.Data().(int)
		fmt.Println(value)
	}

	// Output:
	// 3
	// 2
	// 1
}

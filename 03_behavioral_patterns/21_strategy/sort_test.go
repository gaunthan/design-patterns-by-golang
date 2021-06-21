package strategy

import "fmt"

func ExampleSort() {
	elems := []int{5, 4, 1, 3, 2}

	context := NewSortContext(NewBubbleSortStrategy())
	context.Sort(elems)
	fmt.Println(elems)

	elems = []int{5, 1, 2, 3, 6}

	context.SetStrategy(NewInsertSortStrategy())
	context.Sort(elems)
	fmt.Println(elems)

	// Output:
	// [1 2 3 4 5]
	// [1 2 3 5 6]
}

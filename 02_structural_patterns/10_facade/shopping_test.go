package facade

import (
	"design-patterns-by-golang/02_structural_patterns/10_facade/delivery"
	"fmt"
)

func ExampleOnlineShopping() {
	shopping := NewOnlineShopping()
	outputResult(shopping.Buy("Joe", "apple"))
	outputResult(shopping.Buy("Joe", "orange"))
	outputResult(shopping.Buy("Void", "banana"))

	// Output:
	// ok
	// error: orange out of stock
	// error: destination unreachable
}

func ExampleDelivery() {
	fmt.Println(delivery.SendExpress("Void", "apple"))

	// Output:
	// false
}

func outputResult(ok bool, err error) {
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Printf("error: %v\n", err)
	}
}

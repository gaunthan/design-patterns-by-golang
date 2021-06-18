package composite

import "fmt"

func Example_one() {
	treasureChest := NewChest(TreasureChestType, nil)
	goldenChest := NewChest(ChestOfChestType, treasureChest)
	silverChest := NewChest(ChestOfChestType, goldenChest)
	woodChest := NewChest(ChestOfChestType, silverChest)

	fmt.Println(woodChest.Open())

	// Output:
	// [[[$]]]
}

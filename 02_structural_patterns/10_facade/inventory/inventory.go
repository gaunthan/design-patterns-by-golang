package inventory

type inventorySystem struct {
	stockList map[string]int
}

var system inventorySystem

func init() {
	system = inventorySystem{
		stockList: map[string]int{
			"apple":  1,
			"banana": 2,
		},
	}
}

func TakeOut(item string) bool {
	if system.stockList[item] <= 0 {
		return false
	}
	system.stockList[item]--
	return true
}

func Add(stockList map[string]int) {
	for item, count := range stockList {
		system.stockList[item] += count
	}
}

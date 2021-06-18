package composite

import "fmt"

type Chest interface {
	Store(chest Chest)
	Open() string
}

const (
	TreasureChestType = iota
	ChestOfChestType
)

func NewChest(kind int, holdings Chest) Chest {
	var c Chest
	switch kind {
	case TreasureChestType:
		c = NewTreasureChest(holdings)
	case ChestOfChestType:
		c = NewChestOfChest(holdings)
	}
	return c
}

type chestImpl struct {
	holdings Chest
}

func (p *chestImpl) Store(chest Chest) {
	p.holdings = chest
}

func (p *chestImpl) Open() string {
	return ""
}

type TreasureChest struct {
	chestImpl
}

func (p *TreasureChest) Open() string {
	return "$"
}

func NewTreasureChest(holdings Chest) *TreasureChest {
	return &TreasureChest{}
}

type ChestOfChest struct {
	chestImpl
}

func (p *ChestOfChest) Open() string {
	return fmt.Sprintf("[%s]", p.holdings.Open())
}

func NewChestOfChest(holdings Chest) *ChestOfChest {
	chest := ChestOfChest{}
	chest.Store(holdings)
	return &chest
}

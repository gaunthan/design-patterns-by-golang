package builder

type RoleBuilder interface {
	BuildName()
	BuildSuit()
	BuildWeapon()
	GetResult() *Role
}

type SoldierBuilder struct {
	r Role
}

func (b *SoldierBuilder) BuildName() {
	b.r.Name = "soldier"
}

func (b *SoldierBuilder) BuildSuit() {
	b.r.Suit = "light armour"
}

func (b *SoldierBuilder) BuildWeapon() {
	b.r.Weapon = "sword"
}

func (b *SoldierBuilder) GetResult() *Role {
	r := b.r
	return &r
}

type ArcherBuilder struct {
	r Role
}

func (b *ArcherBuilder) BuildName() {
	b.r.Name = "archer"
}

func (b *ArcherBuilder) BuildSuit() {
	b.r.Suit = "cloth armour"
}

func (b *ArcherBuilder) BuildWeapon() {
	b.r.Weapon = "bow"
}

func (b *ArcherBuilder) GetResult() *Role {
	r := b.r
	return &r
}

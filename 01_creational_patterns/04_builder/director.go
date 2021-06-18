package builder

type Director struct {
	builder RoleBuilder
}

func (d *Director) SetBuilder(builder RoleBuilder) {
	d.builder = builder
}

func (d *Director) Construct() *Role {
	d.builder.BuildName()
	d.builder.BuildSuit()
	d.builder.BuildWeapon()
	return d.builder.GetResult()
}

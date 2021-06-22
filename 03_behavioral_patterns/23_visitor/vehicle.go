package visitor

type Vehicle interface {
	Accept(action VehicleAction)
}

type Motocycle struct{}

func NewMotocycle() *Motocycle {
	return &Motocycle{}
}

func (p *Motocycle) Accept(action VehicleAction) {
	action.VisitMotocycle(p)
}

type Automobile struct{}

func NewAutomobile() *Automobile {
	return &Automobile{}
}

func (p *Automobile) Accept(action VehicleAction) {
	action.VisitAutomobile(p)
}

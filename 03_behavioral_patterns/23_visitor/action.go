package visitor

import "fmt"

type VehicleAction interface {
	VisitMotocycle(motocycle *Motocycle)
	VisitAutomobile(automobile *Automobile)
}

type MoveForwardAction struct{}

func NewMoveForwardAction() *MoveForwardAction {
	return &MoveForwardAction{}
}

func (p *MoveForwardAction) VisitMotocycle(motocycle *Motocycle) {
	fmt.Println("Motocycle is moving forward slowly")
}

func (p *MoveForwardAction) VisitAutomobile(automobile *Automobile) {
	fmt.Println("Automobile is moving forward fastly")
}

type DriftAction struct{}

func NewDriftAction() *DriftAction {
	return &DriftAction{}
}

func (p *DriftAction) VisitMotocycle(motocycle *Motocycle) {
	fmt.Println("Motocycle is drifting carefully")
}

func (p *DriftAction) VisitAutomobile(automobile *Automobile) {
	fmt.Println("Automobile is drifting stably")
}

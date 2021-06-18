package builder

import "fmt"

type Role struct {
	Name   string
	Suit   string
	Weapon string
}

func (r *Role) Intro() string {
	return fmt.Sprintf("%s in %s and %s",
		r.Name, r.Suit, r.Weapon)
}

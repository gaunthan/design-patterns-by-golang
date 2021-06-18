package builder

import "testing"

func TestBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder RoleBuilder
		want    string
	}{
		{
			name:    "SoldierBuilder_test",
			builder: &SoldierBuilder{},
			want:    "soldier in light armour and sword",
		},
		{
			name:    "ArcherBuilder_test",
			builder: &ArcherBuilder{},
			want:    "archer in cloth armour and bow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			director := Director{}
			director.SetBuilder(tt.builder)
			role := director.Construct()

			got := role.Intro()
			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}

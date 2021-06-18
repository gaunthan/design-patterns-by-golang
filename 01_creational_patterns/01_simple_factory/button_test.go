package simplefactory

import "testing"

func TestPress(t *testing.T) {
	tests := []struct {
		Name string
		Type string
		Want string
	}{
		{
			Name: "CircularButton_test",
			Type: "circular",
			Want: "pressing circular button",
		},
		{
			Name: "SquareButton_test",
			Type: "square",
			Want: "pressing square button",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			button := ButtonFactory{}.New(tt.Type)

			got := button.Press()
			if got != tt.Want {
				t.Errorf("%#v got %q want %q",
					tt.Type, got, tt.Want)
			}
		})
	}
}

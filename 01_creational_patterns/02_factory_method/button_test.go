package factorymethod

import "testing"

func TestPress(t *testing.T) {
	tests := []struct {
		Name    string
		Factory ButtonFactory
		Want    string
	}{
		{
			Name:    "CircularButton_test",
			Factory: CircularButtonFactory{},
			Want:    "pressing circular button",
		},
		{
			Name:    "SquareButton_test",
			Factory: SquareButtonFactory{},
			Want:    "pressing square button",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			button := tt.Factory.New()

			got := button.Press()
			if got != tt.Want {
				t.Errorf("got %q want %q", got, tt.Want)
			}
		})
	}
}

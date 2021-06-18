package abstractfactory

import (
	"testing"
)

func TestPress(t *testing.T) {
	tests := []struct {
		Name       string
		WidgetType string
		Factory    WidgetFactory
		Want       string
	}{
		{
			Name:       "CircularButton_test",
			WidgetType: "Button",
			Factory:    CircularWidgetFactory{},
			Want:       "pressing circular button",
		},
		{
			Name:       "SquareButton_test",
			WidgetType: "Button",
			Factory:    SquareWidgetFactory{},
			Want:       "pressing square button",
		},
		{
			Name:       "CircularScrollBar_test",
			WidgetType: "ScrollBar",
			Factory:    CircularWidgetFactory{},
			Want:       "scrolling circular scroll bar",
		},
		{
			Name:       "SquareScrollBar_test",
			WidgetType: "ScrollBar",
			Factory:    SquareWidgetFactory{},
			Want:       "scrolling square scroll bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var got string
			if tt.WidgetType == "Button" {
				button := tt.Factory.NewButton()
				got = button.Press()
			} else {
				scrollBar := tt.Factory.NewScrollBar()
				got = scrollBar.Scroll()
			}

			if got != tt.Want {
				t.Errorf("got %q want %q", got, tt.Want)
			}
		})
	}
}

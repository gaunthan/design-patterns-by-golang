package flyweight

import "testing"

func ExampleDisplayImage() {
	viewer := NewImageViewer("1.png")
	viewer.Display()
	// Output:
	// Display: data of 1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("1.png")
	viewer2 := NewImageViewer("1.png")
	viewer3 := NewImageViewer("2.png")

	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
	if viewer2.ImageFlyweight == viewer3.ImageFlyweight {
		t.Fail()
	}
}

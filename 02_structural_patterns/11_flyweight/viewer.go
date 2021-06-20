package flyweight

import "fmt"

type ImageViewer struct {
	ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	fw := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: fw,
	}
}

func (p *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", p.Data())
}

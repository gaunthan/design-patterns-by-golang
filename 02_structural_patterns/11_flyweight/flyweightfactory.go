package flyweight

import "sync"

type ImageFlyweightFactory struct {
	cache map[string]ImageFlyweight
}

var (
	factory *ImageFlyweightFactory
	once    sync.Once
)

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	once.Do(func() {
		factory = &ImageFlyweightFactory{
			cache: make(map[string]ImageFlyweight),
		}
	})
	return factory
}

func (p *ImageFlyweightFactory) Get(filename string) ImageFlyweight {
	fw := p.cache[filename]
	if fw == nil {
		fw = NewImageFlyweight(filename)
		p.cache[filename] = fw
	}
	return fw
}

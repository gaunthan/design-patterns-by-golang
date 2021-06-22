package iterator

type BidirectionalIterator interface {
	Data() interface{}
	Next() BidirectionalIterator
	Prev() BidirectionalIterator
}

type VectorIterator struct {
	vec   *Vector
	index int
}

func NewVectorIterator(vec *Vector, index int) *VectorIterator {
	return &VectorIterator{
		vec:   vec,
		index: index,
	}
}

func (p *VectorIterator) Data() interface{} {
	return p.vec.values[p.index]
}

func (p *VectorIterator) Next() BidirectionalIterator {
	if p.index >= len(p.vec.values)-1 {
		return nil
	}
	return NewVectorIterator(p.vec, p.index+1)
}

func (p *VectorIterator) Prev() BidirectionalIterator {
	if p.index <= 0 {
		return nil
	}
	return NewVectorIterator(p.vec, p.index-1)
}

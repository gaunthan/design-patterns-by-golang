package iterator

type Vector struct {
	values []interface{}
}

func NewVector() *Vector {
	return &Vector{
		values: []interface{}{},
	}
}

func (p *Vector) Push(v interface{}) {
	p.values = append(p.values, v)
}

func (p *Vector) Pop() interface{} {
	n := len(p.values)
	v := p.values[n-1]
	p.values = p.values[:n-1]
	return v
}

func (p *Vector) Begin() BidirectionalIterator {
	return NewVectorIterator(p, 0)
}

func (p *Vector) End() BidirectionalIterator {
	return nil
}

func (p *Vector) RightBegin() BidirectionalIterator {
	return NewVectorIterator(p, len(p.values)-1)
}

func (p *Vector) RightEnd() BidirectionalIterator {
	return nil
}

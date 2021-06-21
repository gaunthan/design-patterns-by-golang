package strategy

type SortStrategy interface {
	Sort(elems []int)
}

type BubbleSortStrategy struct{}

func NewBubbleSortStrategy() *BubbleSortStrategy {
	return &BubbleSortStrategy{}
}

func (*BubbleSortStrategy) Sort(elems []int) {
	n := len(elems)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if elems[j] > elems[j+1] {
				elems[j], elems[j+1] = elems[j+1], elems[j]
			}
		}
	}
}

type InsertSortStrategy struct{}

func NewInsertSortStrategy() *InsertSortStrategy {
	return &InsertSortStrategy{}
}

func (*InsertSortStrategy) Sort(elems []int) {
	n := len(elems)
	for i := 1; i < n; i++ {
		j := 0
		for ; j < i && elems[j] < elems[i]; j++ {
		}

		v := elems[i]
		for k := i; k > j; k-- {
			elems[k] = elems[k-1]
		}
		elems[j] = v
	}
}

type SortContext struct {
	strategy SortStrategy
}

func NewSortContext(strategy SortStrategy) *SortContext {
	return &SortContext{
		strategy: strategy,
	}
}

func (p *SortContext) Sort(elems []int) {
	p.strategy.Sort(elems)
}

func (p *SortContext) SetStrategy(strategy SortStrategy) {
	p.strategy = strategy
}

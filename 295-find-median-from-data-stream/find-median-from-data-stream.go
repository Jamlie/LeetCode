type LessFunc func(int, int) bool

type Heap struct {
	heap     []int
	lessFunc LessFunc
}

func (a Heap) Len() int           { return len(a.heap) }
func (a Heap) Swap(i, j int)      { a.heap[i], a.heap[j] = a.heap[j], a.heap[i] }
func (a Heap) Less(i, j int) bool { return a.lessFunc(a.heap[i], a.heap[j]) }

func (a Heap) Peek() int { return a.heap[0] }

func (h *Heap) Push(x any) {
	if num, ok := x.(int); ok {
		h.heap = append(h.heap, num)
	}
}

func (h *Heap) Pop() any {
	tempHeap := *h
	size := len(tempHeap.heap)
	last := tempHeap.heap[size-1]
	h.heap = tempHeap.heap[:size-1]
	return last
}

func NewHeap(lessFn LessFunc) *Heap {
	return &Heap{
		heap:     []int{},
		lessFunc: lessFn,
	}
}

type MedianFinder struct {
	smallValues *Heap
	largeValues *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		smallValues: NewHeap(func(i1, i2 int) bool {
			return i1 > i2
		}),
		largeValues: NewHeap(func(i1, i2 int) bool {
			return i1 < i2
		}),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.smallValues, num)

	if mf.smallValues.Len() > mf.largeValues.Len() {
		largeValue := heap.Pop(mf.smallValues).(int)
		heap.Push(mf.largeValues, largeValue)
	}

	if mf.largeValues.Len() > mf.smallValues.Len() {
		smallValue := heap.Pop(mf.largeValues).(int)
		heap.Push(mf.smallValues, smallValue)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.smallValues.Len() > mf.largeValues.Len() {
		return float64(mf.smallValues.Peek())
	}

	if mf.largeValues.Len() > mf.smallValues.Len() {
		return float64(mf.largeValues.Peek())
	}

	minValue := mf.smallValues.Peek()
	maxValue := mf.largeValues.Peek()

	return float64(max(minValue+maxValue)) / 2
}
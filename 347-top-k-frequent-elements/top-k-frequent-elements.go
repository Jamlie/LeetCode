type Element struct {
	num  int
	freq int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }

func (h *MinHeap) Push(x any) {
	switch x.(type) {
	case Element:
		*h = append(*h, x.(Element))
	}
}

func (h *MinHeap) Pop() any {
	tempSlice := *h
	size := len(tempSlice)
	last := tempSlice[size-1]
	*h = tempSlice[:size-1]
	return last
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	h := new(MinHeap)
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, Element{
			num:  num,
			freq: freq,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, 0, k)

	for h.Len() > 0 {
		element := heap.Pop(h)
		switch element.(type) {
		case Element:
			el := element.(Element)
			result = append(result, el.num)
		}
	}

	return result
}
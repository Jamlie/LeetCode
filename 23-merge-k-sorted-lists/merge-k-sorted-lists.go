/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (a MinHeap) Len() int           { return len(a) }
func (a MinHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MinHeap) Less(i, j int) bool { return a[i].Val < a[j].Val }

func (h *MinHeap) Push(x any) {
	if node, ok := x.(*ListNode); ok {
		*h = append(*h, node)
	}
}

func (h *MinHeap) Pop() any {
	tempSlice := *h
	size := len(tempSlice)
	last := tempSlice[size-1]
	*h = tempSlice[:size-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := new(MinHeap)
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	newNode := new(ListNode)
	curr := newNode

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next

		if curr.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return newNode.Next
}
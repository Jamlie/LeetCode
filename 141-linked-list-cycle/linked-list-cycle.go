/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	nodes := make(map[uintptr]bool)
	curr := head

	for curr != nil {
		ptr := uintptr(unsafe.Pointer(curr))

		if nodes[ptr] {
			return true
		}

		nodes[ptr] = true
		curr = curr.Next
	}

	return false    
}
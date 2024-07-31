/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    nodes := make([]*ListNode, 0)
	curr := head

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	n := len(nodes)
	for i := range n / 2 {
		nodes[i].Next = nodes[n-1-i]
		nodes[n-1-i].Next = nodes[i+1]
	}

	nodes[n/2].Next = nil
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	if n < 1 {
		return nil
	}

	curr := head
	nodes := make([]*ListNode, 0)

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	if n > len(nodes) {
		return nil
	}

	index := len(nodes) - n
	if index == 0 {
		head = head.Next
	}

	if index > 0 {
		prevNode := nodes[index-1]

		if prevNode.Next != nil {
			prevNode.Next = prevNode.Next.Next
		}

	}

	return head
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddCurr := head

	if head.Next == nil {
		return oddCurr
	}

	evenCurr := head.Next
	firstEven := evenCurr

	for evenCurr != nil && evenCurr.Next != nil {
		oddCurr.Next = evenCurr.Next
		oddCurr = oddCurr.Next

		evenCurr.Next = oddCurr.Next
		evenCurr = evenCurr.Next
	}

	oddCurr.Next = firstEven

	return head
}
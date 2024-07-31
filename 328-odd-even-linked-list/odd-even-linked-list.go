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
	var evenCurr *ListNode

	if head.Next == nil {
		return oddCurr
	}

	evenCurr = head.Next
	firstEven := evenCurr

	counter := 1

	for oddCurr != nil && evenCurr != nil {
		if counter%2 == 1 {
			if oddCurr.Next != nil {
				oddCurr.Next = oddCurr.Next.Next
			}
			oddCurr = oddCurr.Next
		} else {
			if evenCurr.Next != nil {
				evenCurr.Next = evenCurr.Next.Next
			}
			evenCurr = evenCurr.Next
		}

		counter++
	}

	oddCurrent := head

	for oddCurrent.Next != nil {
		oddCurrent = oddCurrent.Next
	}

	oddCurrent.Next = firstEven

	return head
}
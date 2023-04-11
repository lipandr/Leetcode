package RemoveLinkedListElements

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	var dummy *ListNode
	for {
		if current.Val == val {
			if dummy == nil {
				head = current.Next
			} else {
				dummy.Next = current.Next
			}
		} else {
			dummy = current
		}
		current = current.Next
		if current == nil {
			break
		}
	}
	return head
}

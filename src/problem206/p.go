package problem206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	next := cur.Next
	cur.Next = nil
	for next != nil {
		nextNext := next.Next
		next.Next = cur
		cur = next
		next = nextNext
	}
	return cur
}

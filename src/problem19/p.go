package problem19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		size++
	}

	targetNode := size - n
	if targetNode == 0 {
		return head.Next
	}

	cur = head
	for i := 0; i < targetNode-1; i++ {
		cur = cur.Next
	}
	next := cur.Next
	nextNext := next.Next
	cur.Next = nextNext
	return head
}

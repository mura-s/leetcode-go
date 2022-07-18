package problem876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	middle := length/2 + 1
	for i, node := 1, head; node != nil; i, node = i+1, node.Next {
		if i == middle {
			return node
		}
	}
	return nil
}

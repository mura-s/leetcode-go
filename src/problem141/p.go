package problem141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	foundSet := make(map[*ListNode]bool, 0)
	node := head
	for node != nil {
		if _, ok := foundSet[node]; ok {
			return true
		}
		foundSet[node] = true
		node = node.Next
	}
	return false
}

package problem21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}

		if list1.Val < list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			cur = cur.Next
			list2 = list2.Next
		}
	}
	return dummy.Next
}

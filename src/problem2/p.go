package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p, q, cur := l1, l2, dummyHead
	carry := 0
	for p != nil || q != nil {
		v := carry
		if p != nil {
			v += p.Val
			p = p.Next
		}
		if q != nil {
			v += q.Val
			q = q.Next
		}
		carry = v / 10
		cur.Next = &ListNode{Val: v % 10}
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

func main() {
	fmt.Println(addTwoNumbers(
		&ListNode{Val: 5},
		&ListNode{Val: 5},
	))
}

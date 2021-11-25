package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	preHead := newHead
	cur := head

	for cur != nil {
		if cur.Val == val {
			preHead.Next = cur.Next
		} else {
			preHead = cur
		}
		cur = cur.Next

	}

	return newHead.Next
}

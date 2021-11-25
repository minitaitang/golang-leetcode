package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var fast, slow *ListNode
	fast = head
	slow = head
	step := 0
	for i := 0; i < n; i++ {
		if fast.Next == nil && step < n-1 {
			return head
		}
		fast = fast.Next
		step++
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

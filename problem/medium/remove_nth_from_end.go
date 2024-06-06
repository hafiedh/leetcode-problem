package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: 0, Next: head}
	fast := head
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		start := pre.Next     // 1 -> 2 -> 3 -> 4
		end := pre.Next.Next  // 2 -> 3 -> 4
		pre.Next = end        // 0 -> 2 -> 3 -> 4
		start.Next = end.Next // 1 -> 3 -> 4
		end.Next = start      // 2 -> 1 -> 3 -> 4
		pre = start           // 1 -> 3 -> 4
	}
	return dummy.Next
}

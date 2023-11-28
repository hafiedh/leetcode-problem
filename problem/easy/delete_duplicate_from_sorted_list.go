package problem

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem 1
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

// problem 2
func DeleteDuplicates2(head *ListNode) *ListNode {
	for head != nil && head.Next != nil && head.Val == head.Next.Val {
		for head != nil && head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		head = head.Next
	}
	headFirst, err := json.MarshalIndent(head, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println("First loop: ", string(headFirst))
	if head == nil {
		return nil
	}
	current := head
	for current != nil && current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			for current.Next != nil && current.Next.Next != nil && current.Next.Val == current.Next.Next.Val {
				current.Next = current.Next.Next
			}
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

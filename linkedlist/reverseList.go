package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	newHead := &ListNode{-1, nil}
	for nil != head {
		next := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = next
	}
	return newHead.Next
}

func ReverseByRecursion(head *ListNode) {
	if nil != head {
		ReverseByRecursion(head.Next)
	}

}

func PrintList(head *ListNode) {
	for nil != head {
		fmt.Printf("%d ->", head.Val)
		head = head.Next
	}
}

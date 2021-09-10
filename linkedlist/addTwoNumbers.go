package linkedlist

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

链接：https://leetcode-cn.com/problems/add-two-numbers
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil} // 哨兵节点
	//这里用一个result，只是为了后面返回节点方便，并无他用
	result := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return result.Next
}

func addTwoNumbersByRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{Val: sum, Next: nextNode}
	}
	tempNode := &ListNode{
		Val:  1,
		Next: nil,
	}
	return &ListNode{
		Val:  sum - 10,
		Next: addTwoNumbers(nextNode, tempNode),
	}
}

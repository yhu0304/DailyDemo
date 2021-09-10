package linkedlist

//向链表插入数据，每次都把新增加的结点加到链表尾部
func AddNode(head *ListNode, d int) {
	var newNode *ListNode = &ListNode{d, nil}
	if head == nil {
		head = newNode
		return
	}

	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func DeleteNode(head *ListNode, index int) (b bool) {
	if index < 1 || index > ListLen(head) {
		return
	}
	//删除链表第一个元素
	if 1 == index {
		head = head.Next
		b = true
		return
	}
	var (
		i       = 1
		preNode = head
		curNode = preNode.Next
	)
	for curNode != nil {
		if i == index { //找到删除的结点
			preNode.Next = curNode.Next
			b = true
			return
		}
		preNode = curNode
		curNode = curNode.Next
		i++
	}
	b = true
	return

}

//返回结点的长度
func ListLen(head *ListNode) (length int) {
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}
	return
}

package linkList

import common "tryLeetcode"

//----------------------------------反转链表------------------------------------------

// 反转一个单链表  1->3->5->7->9
func ReverseList(head *common.ListNode) *common.ListNode {
	var prev *common.ListNode
	for head != nil {
		// 保存当前head.Next节点，防止重新赋值后被覆盖
		// 一轮之后状态：nil<-1 2->3->4
		//              prev   head
		temp := head.Next
		head.Next = prev
		// pre 移动
		prev = head
		// head 移动
		head = temp
	}
	return prev
}

// 反转一个单链表， 递归 1->3->5->7
func ReverseListByRecursion(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 反转链表，前n个, 递归 1->3->5->7，3
func ReverseN(head *common.ListNode, n int) *common.ListNode {
	if n == 1 {
		return head
	}
	last := ReverseN(head.Next, n-1)
	// 此时情况类似 (head)1->(3<-5)(last)
	// 							|
	// 							7
	successor := head.Next.Next

	head.Next.Next = head
	head.Next = successor

	return last
}

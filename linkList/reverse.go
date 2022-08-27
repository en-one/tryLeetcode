package linkList

import (
	common "tryLeetcode"
)

//----------------------------------迭代------------------------------------------

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

// 反转单链表， 区间 a，b,返回反转的前部分  1->3->5->7->9
func ReverseN(head, end *common.ListNode) *common.ListNode {

	var prev *common.ListNode
	for head != end {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}

//----------------------------------递归------------------------------------------

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
func ReverseNByRecursion(head *common.ListNode, n int) *common.ListNode {
	if n == 1 {
		return head
	}
	last := ReverseNByRecursion(head.Next, n-1)
	// 此时情况类似 (head)1->(3<-5)(last)
	// 							|
	// 							7
	successor := head.Next.Next

	head.Next.Next = head
	head.Next = successor

	return last
}

// 反转链表， 从m到n个 1->2->3->4->5 【2，4】

//----------------------------------递归+迭代------------------------------------------

// 反转单链表， k个一组， 1->2->3->4->5 2
func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	newHead := ReverseN(head, cur)
	head.Next = reverseKGroup(cur, k)

	return newHead
}

//----------------------------------扩展，回文------------------------------------------

// 寻找中点，反转中点后面的数据  1->2->3->2->1
func isPalindrome(head *common.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 奇数个时，slow还要再往后走一步
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := ReverseList(slow)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

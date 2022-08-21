package linkList

import common "tryLeetcode"

/*
	核心点
		null/nil 异常处理
		dummy node 哑巴节点
		快慢指针
		插入一个节点到排序链表
		从一个链表中移除一个节点
		翻转链表
		合并两个链表
		找到链表的中间节点
*/

// MergeTwoList 合并两个有序链表, 类似归并排序
func MergeTwoList(l1, l2 *common.ListNode) *common.ListNode {

	dummy := &common.ListNode{Val: 0}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dummy.Next
}

// 分隔链表, 并保持原有位置
func Partition(head *common.ListNode, x int) *common.ListNode {
	small := &common.ListNode{Val: 0}
	big := &common.ListNode{Val: 0}
	headSmall := small
	headBig := big
	current := head

	for current != nil {
		if current.Val < x {
			small.Next = current
			small = small.Next
		} else {
			big.Next = current
			big = big.Next
		}
		current = current.Next
	}
	big.Next = nil

	small.Next = headBig.Next
	return headSmall.Next
}

// 合并k个有序链表，关键在于如何寻找多个链表的最小值

// ----------------------------*删除节点----------------------------------------------

// DeleteDuplicatesSaveOne 删除重复节点, 使得每个节点只出现一次
func DeleteDuplicatesSaveOne(head *common.ListNode) *common.ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	return head
}

// deleteDuplicates 删除所有重复节点
func DeleteDuplicatesWithNoOne(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	dummy := &common.ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return dummy.Next
}

//----------------------------------环-----------------------------------------------

// 判断链表是否有环，并返回环节点.同理于判断链表是否有环

// map
func detectCycleForMap(head *common.ListNode) *common.ListNode {
	mcycle := make(map[*common.ListNode]struct{}, 0)

	for head != nil {
		if _, ok := mcycle[head]; ok {
			return head
		}
		mcycle[head] = struct{}{}
		head = head.Next
	}

	return nil
}

// 快慢指针
func detectCycleForQuickSlow(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			// 思路：快慢指针，快慢相遇之后，其中一个回到头，快慢指针步调一致一起移动，相遇点即为入环点
			// 2nb + a = nb + a
			slow = slow.Next
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return nil
}

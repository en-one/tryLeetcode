package linkList

import common "tryLeetcode"

// ----------------------------*删除节点----------------------------------------------

// DeleteDuplicatesSaveOne 删除重复节点, 使得每个节点只出现一次 1->1->2->2->2->3
func DeleteDuplicatesSaveOne(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
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

package tryLeetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// []int{1,2,3} => 1->2->3
func ChangeSliceToListNode(arr []int) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy

	for i := 0; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}

	return dummy.Next
}

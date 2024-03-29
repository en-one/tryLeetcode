package linkList

import (
	"reflect"
	"testing"
	common "tryLeetcode"
)

// 合并两个有序链表
func Test_MergeTwoList(t *testing.T) {

	test1L1 := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 5, Next: nil}}}
	test1L2 := &common.ListNode{Val: -1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 9, Next: &common.ListNode{Val: 16, Next: nil}}}}
	test1Want := &common.ListNode{Val: -1, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 5, Next: &common.ListNode{
		Val: 9, Next: &common.ListNode{Val: 16, Next: nil}}}}}}}

	cases := []struct {
		name  string
		list1 *common.ListNode
		list2 *common.ListNode
		want  *common.ListNode
	}{{"test1", test1L1, test1L2, test1Want}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := MergeTwoList(c.list1, c.list2)
			for get != nil || c.want != nil {
				if get.Val != c.want.Val {
					t.Fatal("failed")
				}
				get = get.Next
				c.want = c.want.Next
			}
		})
	}
}

// 分隔链表
func TestPartition(t *testing.T) {
	type args struct {
		head *common.ListNode
		x    int
	}
	test1 := common.ChangeSliceToListNode([]int{1, 4, 3, 2, 5, 2})
	want1 := common.ChangeSliceToListNode([]int{1, 2, 2, 4, 3, 5})

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args{test1, 3}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 合并k个有序链表
func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*common.ListNode
	}
	test1_1 := common.ChangeSliceToListNode([]int{1, 4, 5})
	test1_2 := common.ChangeSliceToListNode([]int{1, 3, 4})
	test1_3 := common.ChangeSliceToListNode([]int{2, 6})
	want1 := common.ChangeSliceToListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args{[]*common.ListNode{test1_1, test1_2, test1_3}}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 判断链表环-map
func Test_hasCycle_useMap(t *testing.T) {

	head1 := common.ChangeSliceToListNode([]int{-1, 2, 9, 16})

	head2 := common.ChangeSliceToListNode([]int{-1, 2, 9, 16})
	head2.Next.Next = head2.Next

	cases := []struct {
		name string
		head *common.ListNode
		want *common.ListNode
	}{
		{"test-hasCycle-false", head1, nil},
		{"test-hasCycle-false", head2, head2.Next},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := detectCycleForMap(c.head)
			if get != c.want {
				t.Fatal("failed")
			}
		})
	}
}

// 判断链表环-快慢指针
func Test_hasCycle_useQuickSlow(t *testing.T) {

	head1 := common.ChangeSliceToListNode([]int{-1, 2, 9, 16})

	head2 := common.ChangeSliceToListNode([]int{-1, 2, 9, 16})
	head2.Next.Next = head2.Next

	cases := []struct {
		name string
		head *common.ListNode
		want *common.ListNode
	}{
		{"test-hasCycle-false", head1, nil},
		// {"test-hasCycle-true", head2, head2.Next},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := detectCycleForQuickSlow(c.head)
			if get != c.want {
				t.Fatal("failed")
			}
		})
	}
}

// 删除第n个结点
func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *common.ListNode
		n    int
	}

	test1 := common.ChangeSliceToListNode([]int{1, 2, 3, 4, 5})
	want1 := common.ChangeSliceToListNode([]int{1, 2, 3, 5})

	test2 := common.ChangeSliceToListNode([]int{1})

	test3 := common.ChangeSliceToListNode([]int{1, 2})
	want3 := common.ChangeSliceToListNode([]int{1})

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args{head: test1, n: 2}, want1},
		{"test2", args{head: test2, n: 1}, nil},
		{"test3", args{head: test3, n: 1}, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 找链表中点
func Test_middleNode(t *testing.T) {
	type args struct {
		head *common.ListNode
	}

	test1 := common.ChangeSliceToListNode([]int{1, 2, 3, 4, 5})
	want1 := common.ChangeSliceToListNode([]int{3, 4, 5})

	test2 := common.ChangeSliceToListNode([]int{1, 2, 3, 4, 5, 6})
	want2 := common.ChangeSliceToListNode([]int{4, 5, 6})

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args{head: test1}, want1},
		{"test2", args{head: test2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 找链表交点
func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *common.ListNode
		headB *common.ListNode
	}

	args1 := args{
		headA: common.ChangeSliceToListNode([]int{4, 1, 8, 4, 5}),
		headB: common.ChangeSliceToListNode([]int{5, 6, 1, 8, 4, 5}),
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

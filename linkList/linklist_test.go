package linkList

import (
	"reflect"
	"testing"
	common "tryLeetcode"
)

// 合并两个有序链表
func Test_MergeTwoList(t *testing.T) {

	test1L1 := &common.ListNode{1, &common.ListNode{3, &common.ListNode{5, nil}}}
	test1L2 := &common.ListNode{-1, &common.ListNode{2, &common.ListNode{9, &common.ListNode{16, nil}}}}
	test1Want := &common.ListNode{-1, &common.ListNode{1, &common.ListNode{2, &common.ListNode{3, &common.ListNode{5, &common.ListNode{
		9, &common.ListNode{16, nil}}}}}}}

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

// 删除重复节点，保留原始节点
func Test_DeleteDuplicatesSaveOne(t *testing.T) {

	head1 := &common.ListNode{-1, &common.ListNode{-1, &common.ListNode{3, &common.ListNode{3, &common.ListNode{5, &common.ListNode{9, &common.ListNode{16, nil}}}}}}}
	test1Want := &common.ListNode{-1, &common.ListNode{3, &common.ListNode{5, &common.ListNode{9, &common.ListNode{16, nil}}}}}

	cases := []struct {
		name string
		head *common.ListNode
		want *common.ListNode
	}{{"test1", head1, test1Want}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := DeleteDuplicatesSaveOne(c.head)
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

// 删除重复节点,不保留重复数字
func Test_DeleteDuplicatesWithNoOne(t *testing.T) {
	head1 := &common.ListNode{-1, &common.ListNode{-1, &common.ListNode{3, &common.ListNode{3, &common.ListNode{5, &common.ListNode{9, &common.ListNode{16, nil}}}}}}}
	test1Want := &common.ListNode{5, &common.ListNode{9, &common.ListNode{16, nil}}}

	tests := []struct {
		name string
		head *common.ListNode
		want *common.ListNode
	}{{"test1", head1, test1Want}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteDuplicatesWithNoOne(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDuplicatesWithNoOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 判断链表环-map
func Test_hasCycle_useMap(t *testing.T) {

	head1 := &common.ListNode{-1, &common.ListNode{2, &common.ListNode{9, &common.ListNode{16, nil}}}}

	head2 := &common.ListNode{-1, &common.ListNode{2, &common.ListNode{9, &common.ListNode{16, nil}}}}
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

	head1 := &common.ListNode{-1, &common.ListNode{2, &common.ListNode{9, &common.ListNode{16, nil}}}}

	head2 := &common.ListNode{-1, &common.ListNode{2, &common.ListNode{9, &common.ListNode{16, nil}}}}
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
			get := detectCycleForQuickSlow(c.head)
			if get != c.want {
				t.Fatal("failed")
			}
		})
	}
}

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

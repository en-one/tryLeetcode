package linkList

import (
	"reflect"
	"testing"
)

// 合并两个有序链表
func Test_MergeTwoList(t *testing.T) {

	test1L1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	test1L2 := &ListNode{-1, &ListNode{2, &ListNode{9, &ListNode{16, nil}}}}
	test1Want := &ListNode{-1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, &ListNode{
		9, &ListNode{16, nil}}}}}}}

	cases := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
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

	head1 := &ListNode{-1, &ListNode{-1, &ListNode{3, &ListNode{3, &ListNode{5, &ListNode{9, &ListNode{16, nil}}}}}}}
	test1Want := &ListNode{-1, &ListNode{3, &ListNode{5, &ListNode{9, &ListNode{16, nil}}}}}

	cases := []struct {
		name string
		head *ListNode
		want *ListNode
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
	head1 := &ListNode{-1, &ListNode{-1, &ListNode{3, &ListNode{3, &ListNode{5, &ListNode{9, &ListNode{16, nil}}}}}}}
	test1Want := &ListNode{5, &ListNode{9, &ListNode{16, nil}}}

	tests := []struct {
		name string
		head *ListNode
		want *ListNode
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

	head1 := &ListNode{-1, &ListNode{2, &ListNode{9, &ListNode{16, nil}}}}

	head2 := &ListNode{-1, &ListNode{2, &ListNode{9, &ListNode{16, nil}}}}
	head2.Next.Next = head2.Next

	cases := []struct {
		name string
		head *ListNode
		want *ListNode
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

	head1 := &ListNode{-1, &ListNode{2, &ListNode{9, &ListNode{16, nil}}}}

	head2 := &ListNode{-1, &ListNode{2, &ListNode{9, &ListNode{16, nil}}}}
	head2.Next.Next = head2.Next

	cases := []struct {
		name string
		head *ListNode
		want *ListNode
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

// 反转链表
func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	args1 := args{&ListNode{1, &ListNode{3, &ListNode{5, nil}}}}
	want1 := &ListNode{5, &ListNode{3, &ListNode{1, nil}}}

	args2 := args{&ListNode{1, nil}}
	want2 := &ListNode{1, nil}

	args3 := args{nil}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test1", args1, want1},
		{"test2", args2, want2},
		{"empty head", args3, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

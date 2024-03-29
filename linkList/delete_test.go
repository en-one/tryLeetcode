package linkList

import (
	"reflect"
	"testing"
	common "tryLeetcode"
)

// 删除重复节点，保留原始节点
func Test_DeleteDuplicatesSaveOne(t *testing.T) {

	head1 := &common.ListNode{Val: -1, Next: &common.ListNode{Val: -1, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 9, Next: &common.ListNode{Val: 16, Next: nil}}}}}}}
	test1Want := &common.ListNode{Val: -1, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 9, Next: &common.ListNode{Val: 16, Next: nil}}}}}

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
	head1 := &common.ListNode{Val: -1, Next: &common.ListNode{Val: -1, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 9, Next: &common.ListNode{Val: 16, Next: nil}}}}}}}
	test1Want := &common.ListNode{Val: 5, Next: &common.ListNode{Val: 9, Next: &common.ListNode{Val: 16, Next: nil}}}

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

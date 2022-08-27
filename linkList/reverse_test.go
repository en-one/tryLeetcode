package linkList

import (
	"reflect"
	"testing"
	common "tryLeetcode"
)

//----------------------------------迭代------------------------------------------

// 反转链表
func TestReverseList(t *testing.T) {
	type args struct {
		head *common.ListNode
	}

	test1 := common.ChangeSliceToListNode([]int{1, 3, 5})
	want1 := common.ChangeSliceToListNode([]int{5, 3, 1})

	test2 := common.ChangeSliceToListNode([]int{1})
	want2 := common.ChangeSliceToListNode([]int{1})

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args{test1}, want1},
		{"test2", args{test2}, want2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseN(t *testing.T) {
	type args struct {
		head *common.ListNode
		b    *common.ListNode
	}

	link1 := common.ChangeSliceToListNode([]int{1, 2, 3, 5, 7})

	args1 := args{link1, link1.Next.Next}
	want1 := common.ChangeSliceToListNode([]int{2, 1})

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args1, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseN(tt.args.head, tt.args.b)
			for tt.want != nil && got != nil {
				if tt.want.Val != got.Val {
					t.Errorf("ReverseN() = %v, want %v", got.Val, tt.want.Val)
				}
				tt.want = tt.want.Next
				got = got.Next
			}
		})
	}
}

// 回文链表
func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"this is a odd huiwen", args{common.ChangeSliceToListNode([]int{1, 2, 3, 2, 1})}, true},
		{"this is a even huiwen", args{common.ChangeSliceToListNode([]int{1, 2, 2, 1})}, true},
		{"this is not a huiwen", args{common.ChangeSliceToListNode([]int{1, 2, 3, 4, 1})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

//----------------------------------递归------------------------------------------

// 反转链表， 递归
func TestReverseListByRecursion(t *testing.T) {
	type args struct {
		head *common.ListNode
	}

	args1 := args{&common.ListNode{1, &common.ListNode{3, &common.ListNode{5, &common.ListNode{7, nil}}}}}
	want1 := &common.ListNode{7, &common.ListNode{5, &common.ListNode{3, &common.ListNode{1, nil}}}}

	args2 := args{&common.ListNode{1, nil}}
	want2 := &common.ListNode{1, nil}

	args3 := args{nil}

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args1, want1},
		{"test2", args2, want2},
		{"empty head", args3, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseListByRecursion(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseListByRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 反转链表，前n个
func TestReverseNByRecursion(t *testing.T) {
	type args struct {
		head *common.ListNode
		n    int
	}

	args1 := args{&common.ListNode{1, &common.ListNode{3, &common.ListNode{5, &common.ListNode{7, nil}}}}, 3}
	want1 := &common.ListNode{5, &common.ListNode{3, &common.ListNode{1, &common.ListNode{7, nil}}}}

	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args1, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseNByRecursion(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseN() = %v, want %v", got, tt.want)
			}
		})
	}
}

// k个一组反转链表
func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *common.ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{"test1", args{common.ChangeSliceToListNode([]int{1, 2, 3, 4, 5}), 2}, common.ChangeSliceToListNode([]int{2, 1, 4, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			for tt.want != nil && got != nil {
				if tt.want.Val != got.Val {
					t.Errorf("reverseKGroup() = %v, want %v", got.Val, tt.want.Val)
				}
				tt.want = tt.want.Next
				got = got.Next
			}
		})
	}
}

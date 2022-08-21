package linkList

import (
	"reflect"
	"testing"
	common "tryLeetcode"
)

// 反转链表
func TestReverseList(t *testing.T) {
	type args struct {
		head *common.ListNode
	}

	args1 := args{&common.ListNode{1, &common.ListNode{3, &common.ListNode{5, nil}}}}
	want1 := &common.ListNode{5, &common.ListNode{3, &common.ListNode{1, nil}}}

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
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
func TestReverseN(t *testing.T) {
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
			if got := ReverseN(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseN() = %v, want %v", got, tt.want)
			}
		})
	}
}

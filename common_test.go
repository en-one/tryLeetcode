package tryLeetcode

import (
	"reflect"
	"testing"
)

func TestChangeSliceToListNode(t *testing.T) {
	type args struct {
		arr []int
	}

	want1 := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test1", args{[]int{1, 4, 3, 2, 5, 2}}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeSliceToListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeSliceToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package array

import (
	"testing"
)

//
func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 2}}, 2},
		{"test2", args{[]int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 2, 2, 3}, 3}, 2},
		{"test2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

//  移动0
func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"test2", args{[]int{1}}, []int{1}},
		{"test3", args{[]int{0}}, []int{0}},
		{"test3", args{[]int{0, 1}}, []int{1, 0}},
		{"test3", args{[]int{1, 0}}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)

		})
	}
}

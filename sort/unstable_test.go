package sort

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{8, 1, 19, 5, 66}}, []int{1, 5, 8, 19, 66}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuick(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{8, 1, 19, 5, 66}}, []int{1, 5, 8, 19, 66}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quick(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{8, 1, 19, 5, 66}}, []int{1, 5, 8, 19, 66}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Heap(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShell(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{8, 1, 19, 5, 66}}, []int{1, 5, 8, 19, 66}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shell(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shell() = %v, want %v", got, tt.want)
			}
		})
	}
}

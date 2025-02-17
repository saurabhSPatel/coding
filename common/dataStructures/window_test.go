package common

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"tt1",
			args{
				[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3,
			},
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			"tt2",
			args{
				[]int{1}, 1,
			},
			[]int{1},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"tt1",
			args{
				7,
				[]int{2, 3, 1, 2, 4, 3},
			},
			2,
		},
		{
			"tt2",
			args{
				11,
				[]int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

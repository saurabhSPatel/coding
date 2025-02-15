package common

import (
	"fmt"
	"testing"
)

func Test_dfsTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tt1",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dfsTraversal(tt.args.root)
			fmt.Println(got)
		})
	}
}

func Test_bfsTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"tt1",
			args{
				root: &TreeNode{
					1,
					nil, nil,
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bfsTraversal(tt.args.root)
		})
	}
}

package common

import "testing"

func Test_graphDfs(t *testing.T) {
	type args struct {
		node    *GraphNode
		visited map[*GraphNode]bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graphDfs(tt.args.node, tt.args.visited)
		})
	}
}

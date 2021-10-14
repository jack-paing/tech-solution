package algo_test

import (
	"spenmo-tech/algo"
	"testing"
)

func Test_sequenceExists(t *testing.T) {
	type args struct {
		main []int
		seq  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test for false", args: args{main: []int{20, 7, 8, 10, 5, 6}, seq: []int{1, 4}}, want: false},
		{name: "Test for false", args: args{main: []int{20, 7, 8, 10, 5, 6}, seq: []int{8, 7}}, want: false},
		{name: "Test for false", args: args{main: []int{20, 7, 8, 10, 5, 6}, seq: []int{8, 9}}, want: false},
		{name: "Test for false", args: args{main: []int{20, 7, 8, 10, 5, 6}, seq: []int{1, 2}}, want: false},
		{name: "Test for true", args: args{main: []int{20, 7, 8, 10, 5, 6}, seq: []int{7, 8}}, want: true},
		{name: "Test for true", args: args{main: []int{20, 7, 8, 10, 5, 6}, seq: []int{5, 6}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algo.SequenceExists(tt.args.main, tt.args.seq); got != tt.want {
				t.Errorf("sequenceExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_Find(t *testing.T) {
	type fields struct {
		Val   int
		Left  *algo.TreeNode
		Right *algo.TreeNode
	}
	tn := fields{
		Val: 6,
		Left: &algo.TreeNode{
			Val: 7,
			Left: &algo.TreeNode{
				Val: 10,
			},
			Right: &algo.TreeNode{
				Val: 8,
			},
		},
		Right: &algo.TreeNode{
			Val: 100,
			Right: &algo.TreeNode{
				Val: 14,
			},
		},
	}
	tests := []struct {
		name   string
		fields fields
		key    int
		want   bool
	}{
		{name: "Test for False", fields: tn, key: 91, want: false},
		{name: "Test for False", fields: tn, key: 17, want: false},
		{name: "Test for False", fields: tn, key: -8, want: false},
		{name: "Test for False", fields: tn, key: 111, want: false},
		{name: "Test for True", fields: tn, key: 6, want: true},
		{name: "Test for True", fields: tn, key: 7, want: true},
		{name: "Test for True", fields: tn, key: 8, want: true},
		{name: "Test for True", fields: tn, key: 10, want: true},
		{name: "Test for True", fields: tn, key: 14, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := &algo.TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := tn.Find(tt.key); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

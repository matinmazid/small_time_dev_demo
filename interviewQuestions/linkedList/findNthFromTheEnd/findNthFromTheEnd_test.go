package main

import (
	"testing"
)

var nodeValues = []string{"first", "second", "third", "fourth", "fifth"}

func generateTestData() *node {

	ptr := &node{}
	headPtr := ptr
	for index, value := range nodeValues {
		ptr.Value = value
		if index+1 <= len(nodeValues) {
			nextNode := node{}
			ptr.Ptr = &nextNode
			ptr = &nextNode
		}
	}
	return headPtr
}

func Test_expectedSolution(t *testing.T) {
	type args struct {
		headPtr *node
		nth     int
	}

	headPtr := generateTestData()
	tests := []struct {
		name   string
		args   args
		result *node
	}{
		{"we should  find the fourth", args{headPtr, 2}, &node{nil, "fourth"}},
		{"we should  find the second", args{headPtr, 4}, &node{nil, "second"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := expectedSolution(tt.args.headPtr, tt.args.nth)
			if actual.Value != tt.result.Value {
				t.Errorf("failed idnt get %s" + tt.result.Value)
			}
		})
	}
}

func Test_mySolution(t *testing.T) {

	type args struct {
		headPtr *node
		nth     int
	}

	headPtr := generateTestData()
	tests := []struct {
		name   string
		args   args
		result *node
	}{
		{"we should  find the fourth", args{headPtr, 2}, &node{nil, "fourth"}},
		{"we should  find the second", args{headPtr, 4}, &node{nil, "second"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mySolution(tt.args.headPtr, tt.args.nth)
		})
	}
}

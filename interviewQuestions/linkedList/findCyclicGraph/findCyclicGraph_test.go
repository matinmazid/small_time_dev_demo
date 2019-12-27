package main

import (
	"testing"
)

func Test_generateGraph(t *testing.T) {
	got := generateGraph("a,b,c")

	if got.Value != "a" {
		t.Errorf("didnt get a")
	}

	if got.NextNode.Value != "b" {
		t.Errorf("didnt get b")
	}

	if got.NextNode.NextNode.Value != "c" {
		t.Error("didnt get c")
	}
}

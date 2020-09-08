package main

import (
	"testing"
)

func Test_generateGraph(t *testing.T) {
	got := generateGraph("a:b,b:c,c")

	if got.Name != "a" {
		t.Errorf("didnt get a")
	}

	if got.NextNode.Name != "b" {
		t.Errorf("didnt get b")
	}

	if got.NextNode.NextNode.Name != "c" {
		t.Errorf("didnt get c")
	}

	if got.NextNode.NextNode.NextNode != nil {
		t.Errorf("got an extra dmmy node at the end")
	}
}

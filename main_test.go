package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	h := heap{}
	h.Insert(11)
	h.Insert(3)
	h.Insert(4)
	h.Insert(8)
	h.Insert(9)
	h.Insert(7)
	h.Insert(10)
	h.Insert(19)
	v, _ := h.PopMin()
	if v != 3 {
		t.Errorf("got %d, want %d", v, 3)
	}
	v, _ = h.PopMin()
	if v != 4 {
		t.Errorf("got %d, want %d", v, 4)
	}
	h.Insert(1)
	v, _ = h.PopMin()
	if v != 1 {
		t.Errorf("got %d, want %d", v, 1)
	}

}

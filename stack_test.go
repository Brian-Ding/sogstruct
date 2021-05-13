package sogstruct

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := newStack()
	s.Push("1")

	if want, got := "1", s.Pop(); got != want {
		t.Errorf("Stack does not insert/remove the 1")
	}

	s.Push("1")
	s.Push("2")
	s.Push("3")

	if want, got := "3", s.Pop(); got != want {
		t.Errorf("Stack does not insert/remove the 3")
	}
	if want, got := "2", s.Pop(); got != want {
		t.Errorf("Stack does not insert/remove the 2")
	}
	if want, got := "1", s.Pop(); got != want {
		t.Errorf("Stack does not insert/remove the 1")
	}

	wantCount := 0
	if gotCount := len(s.values); gotCount != wantCount {
		t.Errorf("Stack still has values left after poping the last value")
	}
}

func TestStackClear(t *testing.T) {
	s := newStack()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Push("4")
	s.Clear()

	want := 0
	if got := len(s.values); got != want {
		t.Errorf("Clear() has %d objects left, want %d", len(s.values), want)
	}
}

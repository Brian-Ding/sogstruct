package sogstruct

import (
	"testing"
)

func TestLinkedListAddAfter(t *testing.T) {
	l := NewLinkedList()
	l.AddFirstValue("1")

	l.AddAfter(l.head, NewLinkedListNode("2"))
	if want, got := "2", l.head.next.value; want != got {
		t.Errorf("\"2\" is not added after \"1\"")
	}

	l.AddAfter(l.head.next, NewLinkedListNode("3"))
	if want, got := "3", l.head.next.next.value; want != got {
		t.Errorf("\"3\" is not added after \"2\"")
	}
}

func TestLinkedListAddAfterValue(t *testing.T) {
	l := NewLinkedList()
	l.AddFirstValue("1")

	l.AddAfterValue(l.head, "2")
	if want, got := "2", l.head.next.value; want != got {
		t.Errorf("\"2\" is not added after \"1\"")
	}

	l.AddAfterValue(l.head.next, "3")
	if want, got := "3", l.head.next.next.value; want != got {
		t.Errorf("\"3\" is not added after \"2\"")
	}
}

func TestLinkedListAddBefore(t *testing.T) {
	l := NewLinkedList()
	l.AddFirstValue("1")

	l.AddBefore(l.head, NewLinkedListNode("2"))
	if want, got := "2", l.head.value; want != got {
		t.Errorf("\"2\" is not added before \"1\"")
	}

	l.AddBefore(l.head, NewLinkedListNode("3"))
	if want, got := "3", l.head.value; want != got {
		t.Errorf("\"3\" is not added before \"2\"")
	}
}

func TestLinkedListAddBeforeValue(t *testing.T) {
	l := NewLinkedList()
	l.AddFirstValue("1")

	l.AddBeforeValue(l.head, "2")
	if want, got := "2", l.head.value; want != got {
		t.Errorf("\"2\" is not added before \"1\"")
	}

	l.AddBeforeValue(l.head, "3")
	if want, got := "3", l.head.value; want != got {
		t.Errorf("\"3\" is not added before \"2\"")
	}
}

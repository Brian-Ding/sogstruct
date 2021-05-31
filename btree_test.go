package sogstruct

import (
	"testing"
)

func TestBTree(t *testing.T) {
	tree := newBTree(4)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(80)
	tree.Insert(30)
	tree.Insert(35)

	if want, got := 30, tree.root.keys[0]; got != want {
		t.Errorf("key of btree's root is not [30]")
	}

	if want, got := 40, tree.root.keys[1]; got != want {
		t.Errorf("key of btree's root is not [40]")
	}

	if want, got := 70, tree.root.keys[2]; got != want {
		t.Errorf("key of btree's root is not [70]")
	}

	if want, got := 10, tree.root.children[0].keys[0]; got != want {
		t.Errorf("key of btree's children[0] is not [10]")
	}

	if want, got := 20, tree.root.children[0].keys[1]; got != want {
		t.Errorf("key of btree's children[0] is not [20]")
	}

	if want, got := 35, tree.root.children[1].keys[0]; got != want {
		t.Errorf("key of btree's children[1] is not [35]")
	}

	if want, got := 50, tree.root.children[2].keys[0]; got != want {
		t.Errorf("key of btree's children[2] is not [50]")
	}

	if want, got := 60, tree.root.children[2].keys[1]; got != want {
		t.Errorf("key of btree's children[2] is not [60]")
	}

	if want, got := 80, tree.root.children[3].keys[0]; got != want {
		t.Errorf("key of btree's children[3] is not [80]")
	}
}

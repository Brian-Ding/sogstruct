package sogstruct

import (
	"math"
)

type btreenode struct {
	m        int
	keys     []int
	children []*btreenode
	parent   *btreenode
}

// 1. only create the next node, if current node have at least [m/2] children, m the maximum number of allowed children
// 2. root can have 2children: means only 1 key
// 3. all leaf nodes at the same level
// 4. creation process is bottom up
type btree struct {
	m    int
	root *btreenode
}

// m children, m-1 keys
func newBTree(m int) *btree {
	return &btree{m: m, root: &btreenode{m: m, keys: make([]int, 0), children: make([]*btreenode, 0)}}
}

func (tree *btree) Insert(value int) {
	tree.root.insert(value, tree)
}

func (node *btreenode) insert(value int, tree *btree) {
	if len(node.children) == 0 {
		// node is leaf, otherwise these is no chance the node has no children
		if !node.addkey(value) {
			node.propagate(tree)
		}

		return
	}

	// node is not leaf, it must have at least 2 children
	for i := 0; i < len(node.keys); i++ {
		if value < node.keys[i] {
			// always insert into its child
			node.children[i].insert(value, tree)

			return
		}
	}

	// value is larger than every key
	node.children[len(node.keys)].insert(value, tree)
}

func (node *btreenode) propagate(tree *btree) {
	// this node has m keys, which is larger than maximum allowed m-1
	boundary := int(math.Ceil(float64(node.m) / 2))
	leftkeys := node.keys[:boundary]
	upkey := node.keys[boundary]
	rightkeys := node.keys[boundary+1:]
	node.keys = leftkeys

	if node.parent != nil {
		// this node is not root
		// its parent has at least 1 key
		// its parent has at least 2 children
		if !node.parent.addkey(upkey) {
			node.parent.propagate(tree)
		}

		newnode := &btreenode{m: node.m, keys: make([]int, 0), children: make([]*btreenode, 0)}
		for _, v := range rightkeys {
			// no chance the newnode will be full at this stage
			// since only add partial keys from another full node
			newnode.addkey(v)
		}
		boundaryparent := len(node.parent.keys)
		for i, v := range node.parent.keys {
			if rightkeys[0] < v {
				boundaryparent = i
				break
			}
		}
		temp := make([]btreenode, 0)
		for _, v := range node.parent.children[boundaryparent:] {
			temp = append(temp, *v)
		}
		node.parent.children = node.parent.children[:boundaryparent]
		node.parent.children = append(node.parent.children, newnode)
		for _, v := range temp {
			n := &btreenode{m: v.m, keys: v.keys, children: v.children, parent: v.parent}
			node.parent.children = append(node.parent.children, n)
		}
		newnode.parent = node.parent

		return
	}

	// this node is root
	tree.root = &btreenode{m: node.m, keys: make([]int, 0), children: make([]*btreenode, 0)}
	tree.root.addkey(upkey)
	// 1st children
	tree.root.children = append(tree.root.children, node)
	node.parent = tree.root

	// 2nd children
	newnode := &btreenode{m: node.m, keys: make([]int, 0), children: make([]*btreenode, 0)}
	for _, v := range rightkeys {
		// no chance the newnode will be full at this stage
		// since only add partial keys from another full node
		newnode.addkey(v)
	}
	tree.root.children = append(tree.root.children, newnode)
	newnode.parent = tree.root
}

func (node *btreenode) addkey(value int) bool {
	if len(node.keys) == 0 {
		// this node has no keys
		// add the key directly
		node.keys = append(node.keys, value)

		return true
	}

	temp := make([]int, 0)
	for i := 0; i < len(node.keys); i++ {
		if node.keys[i] <= value {
			temp = append(temp, node.keys[i])
		} else {
			temp = append(temp, value)
			temp = append(temp, node.keys[i:]...)
			node.keys = temp

			return len(node.keys) <= node.m-1
		}
	}

	temp = append(temp, value)
	node.keys = temp

	return len(node.keys) <= node.m-1
}

func (tree *btree) Delete(value int) {
	// find where is the value in the tree
	node := tree.root.search(value)
	if node == nil {
		// value not in the tree
		return
	}

	node.delete(value, tree)
}

func (node *btreenode) delete(value int, tree *btree) {
	if len(node.children) == 0 {
		// leaf node
		// remove the key directory
		temp := node.keys
		node.keys = make([]int, 0, len(temp)-1)
		for _, v := range temp {
			if v != value {
				node.keys = append(node.keys, v)
			}
		}
	} else {
		// internal node
		var index int
		for i, v := range node.keys {
			if v == value {
				index = i
				break
			}
		}

		// child before the key has more than minimum keys
		// remove the key, and fill in child's last key
		if len(node.children[index].keys) > int(math.Ceil(float64(node.m)/2))-1 {
			node.keys[index] = node.children[index].keys[len(node.children[index].keys)-1]
			node.children[index].keys = node.children[index].keys[:len(node.children[index].keys)-1]
		} else if len(node.children[index+1].keys) > int(math.Ceil(float64(node.m)/2))-1 {
			// child after the key has more than minimum keys
			// remove the key, and fill in child's first key
			node.keys[index] = node.children[index+1].keys[0]
			node.children[index+1].keys = node.children[index+1].keys[1:]
		} else {
			tempKeys := node.keys
			node.keys = make([]int, 0, len(tempKeys)-1)
			for _, v := range tempKeys {
				if v != value {
					node.keys = append(node.keys, v)
				}
			}

			tempChildren := node.children
			node.children = make([]*btreenode, 0, len(tempChildren)-1)
			for i := 0; i < len(tempChildren); i++ {
				child := &btreenode{m: node.children[i].m, keys: node.children[i].keys, children: node.children[i].children, parent: node.children[i].parent}
				node.children = append(node.children, child)
				if i == index {
					i++
					child = &btreenode{m: node.children[i].m, keys: node.children[i].keys, children: node.children[i].children, parent: node.children[i].parent}
					node.children = append(node.children, child)
				}
			}
		}
	}

	if len(node.keys) < int(math.Ceil(float64(node.m)/2))-1 {
		// number of keys less than allowed minimum
		// get one key from its parent
		node.stealpkey(tree)
	}
}

func (node *btreenode) stealpkey(tree *btree) {
	// root
	if node.parent == nil {
		if len(node.keys) == 0 {
			tree.root.keys = node.children[0].keys
			tree.root.keys = append(tree.root.keys, node.children[1].keys...)

			tree.root.children = node.children[0].children
			tree.root.children = append(tree.root.children, node.children[1].children...)
		} else {
			merged := false
			for i, v := range node.keys {
				if merged {
					node.children[i].keys = node.children[i+1].keys
					node.children[i].children = node.children[i+1].children
				} else if node.children[i+1].keys[0] < v {
					node.children[i].keys = append(node.children[i].keys, node.children[i+1].keys...)
					node.children[i].children = append(node.children[i].children, node.children[i+1].children...)

					merged = true
				}
			}

			if !merged {
				node.children[len(node.children)-2].keys = append(node.children[len(node.children)-2].keys, node.children[len(node.children)-1].keys...)
				node.children[len(node.children)-2].children = append(node.children[len(node.children)-2].children, node.children[len(node.children)-1].children...)
			}

			node.children = node.children[:len(node.keys)+1]
		}
		return
	}

	var index int
	for i, v := range node.parent.children {
		if v == node {
			index = i
			break
		}
	}

	// the key stolen from parent
	var key int
	if index == len(node.parent.keys) {
		// if the last children, get the last key from parent
		index--
		temp := node.parent.keys
		key = temp[index]
		node.parent.keys = make([]int, 0, len(temp)-1)
		for i, v := range temp {
			if i != index {
				node.parent.keys = append(node.parent.keys, v)
			}
		}
	}
	node.addkey(key)
	node.parent.stealpkey(tree)
}

func (node *btreenode) search(value int) *btreenode {
	if len(node.children) == 0 {
		// no children, leaf node
		for _, v := range node.keys {
			if v == value {
				return node
			}
		}

		return nil
	}

	// not leaf node
	for i, v := range node.keys {
		if value < v {
			return node.children[i].search(value)
		} else if value == v {
			return node
		}
	}

	return node.children[len(node.keys)].search(value)
}

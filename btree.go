package sogstruct

import "math"

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

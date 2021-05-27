package sogstruct

type bpTreeNode struct {
	value int
}

type bpTree struct {
	children []*bpTreeNode
}

func newBPTreeNode(value int) *bpTreeNode {
	return &bpTreeNode{value: value}
}

func NewBPTree() *bpTree {
	return &bpTree{children: make([]*bpTreeNode, 0)}
}

func (tree *bpTree) Add(value int) {
	tree.children = append(tree.children, newBPTreeNode(value))
}

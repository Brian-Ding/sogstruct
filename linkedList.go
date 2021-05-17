package sogstruct

type linkedListNode struct {
	next  *linkedListNode
	value string
}

type linkedList struct {
	head *linkedListNode
	// tail *LinkedListNode
}

func newLinkedListNode(value string) *linkedListNode {
	return &linkedListNode{value: value}
}

func newLinkedList() *linkedList {
	return &linkedList{}
}

func (list *linkedList) Find(value string) *linkedListNode {
	if list == nil {
		return nil
	}

	node := list.head
	for {
		if node == nil {
			return nil
		}

		if node.value == value {
			return node
		}

		node = list.head.next
	}
}

func (list *linkedList) AddAfter(node *linkedListNode, newNode *linkedListNode) {
	if list == nil || node == nil {
		return
	}

	listNode := list.head
	for {
		if listNode == nil {
			return
		}

		if listNode == node {
			next := listNode.next
			listNode.next = newNode
			if newNode != nil {
				newNode.next = next
			}

			break
		}

		listNode = listNode.next
	}
}

func (list *linkedList) AddAfterValue(node *linkedListNode, value string) {
	list.AddAfter(node, newLinkedListNode(value))
}

func (list *linkedList) AddBefore(node *linkedListNode, newNode *linkedListNode) {
	if list == nil || node == nil {
		return
	}

	var lastNode *linkedListNode
	listNode := list.head
	for {
		if listNode == nil {
			return
		}

		if listNode == node {
			if lastNode == nil {
				list.head = newNode
			} else {
				lastNode.next = newNode
			}
			if newNode != nil {
				newNode.next = listNode
			}

			break
		}

		lastNode = listNode
		listNode = listNode.next
	}
}

func (list *linkedList) AddBeforeValue(node *linkedListNode, value string) {
	list.AddBefore(node, newLinkedListNode(value))
}

func (list *linkedList) AddFirst(newNode *linkedListNode) {
	if list == nil {
		return
	}

	oldHead := list.head
	list.head = newNode
	if newNode != nil {
		list.head.next = oldHead
	}
}

func (list *linkedList) AddFirstValue(value string) {
	list.AddFirst(newLinkedListNode(value))
}

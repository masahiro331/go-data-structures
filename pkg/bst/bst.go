package bst

import "fmt"

type Node struct {
	leftPtr  *Node
	rightPtr *Node
	Value    int
}

func (n *Node) Insert(v int) {
	if n.less(v) {
		if n.rightPtr == nil {
			n.rightPtr = new(v)
			return
		}
		n.rightPtr.Insert(v)
	} else {
		if n.leftPtr == nil {
			n.leftPtr = new(v)
			return
		}
		n.leftPtr.Insert(v)
	}
}

func (n *Node) equal(v int) bool {
	return n.Value == v
}

func (n *Node) Remove(v int) *Node {

	if n.less(v) {
		if n.rightPtr != nil {
			n.rightPtr = n.rightPtr.Remove(v)
		}
	} else if n.equal(v) {
		if n.leftPtr == nil {
			return n.rightPtr
		} else if n.rightPtr == nil {
			return n.leftPtr
		}

		tmp := n.leftPtr.max()
		n.Value = tmp.Value
		tmp.leftPtr.Remove(tmp.Value)
	} else {
		if n.leftPtr != nil {
			n.leftPtr = n.leftPtr.Remove(v)
		}
	}

	return n
}

func (n *Node) max() *Node {
	if n.rightPtr == nil {
		return n
	} else {
		return n.rightPtr.max()
	}
}

func (n *Node) Search(v int) bool {
	if n.equal(v) {
		return true
	}

	if n.less(v) {
		if n.rightPtr != nil {
			return n.rightPtr.Search(v)
		}
	} else {
		if n.leftPtr != nil {
			return n.leftPtr.Search(v)
		}
	}
	return false
}

func (n *Node) Inorder() {
	if n.leftPtr != nil {
		n.leftPtr.Inorder()
	}
	fmt.Println(n.Value)
	if n.rightPtr != nil {
		n.rightPtr.Inorder()
	}
}

func new(v int) *Node {
	return &Node{
		Value: v,
	}
}

func (n *Node) less(v int) bool {
	return n.Value < v
}

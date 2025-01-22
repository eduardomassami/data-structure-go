package binarytree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) Add(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Add(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Add(value)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if n.Value == value {
		return true
	} else if n.Value < value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

func (n *Node) PrintInOrder(node *Node, level int) {
	if node == nil {
		return
	}

	n.PrintInOrder(node.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.Value)
	n.PrintInOrder(node.Left, level+1)
}

func (n *Node) InvertBTree() {
	if n == nil {
		return
	}

	n.Left, n.Right = n.Right, n.Left

	n.Left.InvertBTree()
	n.Right.InvertBTree()
}

func (n *Node) Heigth() int {
	if n == nil {
		return -1
	}

	heigthLegth := n.Left.Heigth()
	heightRigth := n.Right.Heigth()

	if heigthLegth > heightRigth {
		return heigthLegth + 1
	}
	return heightRigth + 1
}

func (n *Node) Smallest() int {
	if n == nil {
		panic("Btree empty")
	}

	v := n
	for v.Left != nil {
		v = v.Left
	}

	return v.Value
}

func (n *Node) Greatest() int {
	if n == nil {
		panic("Btree empty")
	}

	v := n
	for v.Right != nil {
		v = v.Right
	}

	return v.Value
}

func (n *Node) IsBST(min, max *int) bool {
  if n == nil {
    return true
  }

  if (min != nil && n.Value <= *min) || (max != nil && n.Value >= *max) {
    return false
  }

  return n.Left.IsBST(min, &n.Value) && n.Right.IsBST(&n.Value, max)
}

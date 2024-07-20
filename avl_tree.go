package main

import "fmt"

type AVLTree struct {
	root *Node
}

func NewALVTree() *AVLTree {
	return &AVLTree{root: nil}
}

func (a *AVLTree) Print() {
	printTree(a.root, "", true)
}

func (a *AVLTree) Insert(val int) {
	a.root = insert(a.root, val)
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
	} else {
		return root
	}

	bf := balanceFactor(root)
	if isImbalance(bf) {
		imbalanceType := imbalanceType(root, bf, val)
		root = rotate(root, imbalanceType)
	}

	return root
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if rightHeight > leftHeight {
		return rightHeight + 1
	}

	return leftHeight + 1
}

func balanceFactor(root *Node) int {
	return height(root.Left) - height(root.Right)
}

func isImbalance(balanceFactor int) bool {
	return balanceFactor < -1 || balanceFactor > 1
}

func imbalanceType(root *Node, balanceFactor int, insertedVal int) Imbalance {
	var imbalance Imbalance
	if balanceFactor > 1 && insertedVal < root.Left.Val {
		imbalance = LeftLeftImbalance
	} else if balanceFactor > 1 && insertedVal > root.Left.Val {
		imbalance = LeftRightImbalance
	} else if balanceFactor < -1 && insertedVal > root.Right.Val {
		imbalance = RightRightImbalance
	} else if balanceFactor < -1 && insertedVal < root.Right.Val {
		imbalance = RightLeftImbalance
	}

	return imbalance
}

func rotate(root *Node, imbalanceType Imbalance) *Node {
	switch imbalanceType {
	case LeftLeftImbalance:
		{
			root = rightRotate(root)
		}
	case RightRightImbalance:
		{
			root = leftRotate(root)
		}
	case LeftRightImbalance:
		{
			root.Left = leftRotate(root.Left)
			root = rightRotate(root)
		}
	case RightLeftImbalance:
		{
			root.Right = rightRotate(root.Right)
			root = leftRotate(root)
		}
	}

	return root
}

func rightRotate(root *Node) *Node {
	n1 := root
	n2 := n1.Left
	t1 := n2.Right

	n2.Right = n1
	n1.Left = t1

	return n2
}

func leftRotate(root *Node) *Node {
	n1 := root
	n2 := root.Right
	t1 := n2.Left

	n2.Left = n1
	n1.Right = t1

	return n2
}

func printTree(node *Node, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	if isLeft {
		fmt.Printf("%s└── %d\n", prefix, node.Val)
		prefix += "    "
	} else {
		fmt.Printf("%s┌── %d\n", prefix, node.Val)
		prefix += "│   "
	}

	if node.Left != nil || node.Right != nil {
		printTree(node.Right, prefix, false)
		printTree(node.Left, prefix, true)
	}
}

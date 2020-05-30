package binary_search_tree

import "fmt"

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

func TraverseNode(node *BSTNode, f func(node *BSTNode)) {
	f(node)
	if node.Left != nil {
		TraverseNode(node.Left, f)
	}
	if node.Right != nil {
		TraverseNode(node.Right, f)
	}
	return
}

func PrintTree(node *BSTNode) {
	if node.Left != nil {
		PrintTree(node.Left)
	}
	if node.Right != nil {
		PrintTree(node.Right)
	}
	fmt.Println(node.Value)
}

func (bst *BSTree) Print() {
	PrintTree(bst.root)
}

func BalancedNodeInsertTo(node *BSTNode, value int) {
	if node.Value > value {
		if node.Left == nil {
			node.Left = &BSTNode{
				Value: value,
			}
		} else {
			BalancedNodeInsertTo(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BSTNode{
				Value: value,
			}

		} else {
			BalancedNodeInsertTo(node.Right, value)
		}
	}
}

func BinarySearchValue(node *BSTNode, value int) bool {
	found := false
	if node.Value == value {
		return true
	}
	if node.Value > value {
		if node.Left != nil {
			found = BinarySearchValue(node.Left, value)
		} else {
			return found
		}
	}
	if node.Value < value {
		if node.Right != nil {
			found = BinarySearchValue(node.Right, value)
		} else {
			return found
		}
	}
	return found
}

type BSTree struct {
	root   *BSTNode
	Height int
}

func (bst *BSTree) Root() *BSTNode {
	return bst.root
}

func NewBSTree(value int) *BSTree {
	newNode := &BSTNode{
		Value: value,
	}
	return &BSTree{
		root:   newNode,
		Height: 0,
	}
}
func (bst *BSTree) Insert(value int) {
	BalancedNodeInsertTo(bst.Root(), value)
}

func (bst *BSTree) Search(value int) bool {
	return BinarySearchValue(bst.root, value)
}

package binary_search_tree

import (
	"../binary_search_tree"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestNewBSTree(t *testing.T) {
	bst := binary_search_tree.NewBSTree(50)
	assert.EqualValues(t, 50, bst.Root().Value)
}

func TestTraverseNode(t *testing.T) {
	bst := binary_search_tree.NewBSTree(50)
	out := strings.Builder{}
	f := func(node *binary_search_tree.BSTNode) {
		out.WriteString(strconv.Itoa(node.Value))
	}
	binary_search_tree.TraverseNode(bst.Root(), f)
	assert.EqualValues(t, "50", out.String())
}

func TestBalancedNodeInsertTo(t *testing.T) {
	bst := binary_search_tree.NewBSTree(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(35)
	bst.Insert(55)
	bst.Insert(65)
	out := strings.Builder{}
	f := func(node *binary_search_tree.BSTNode) {
		out.WriteString(strconv.Itoa(node.Value))
	}
	binary_search_tree.TraverseNode(bst.Root(), f)
	assert.EqualValues(t, "505565303520", out.String())
	bst.Print()
}

func TestBinarySearch(t *testing.T) {
	bst := binary_search_tree.NewBSTree(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(35)
	bst.Insert(55)
	bst.Insert(65)
	assert.True(t, bst.Search(35))

}

func TestBinarySearchValueNotFound(t *testing.T) {
	bst := binary_search_tree.NewBSTree(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(35)
	bst.Insert(55)
	bst.Insert(65)
	assert.False(t, bst.Search(100))
}

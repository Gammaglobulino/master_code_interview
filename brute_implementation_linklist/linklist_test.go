package brute_implementation_linklist

import (
	"../brute_implementation_linklist"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateList(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	assert.EqualValues(t, 1, linkList.Size())
	log.Println(linkList)
}
func TestShowList(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30).Append(40)
	assert.EqualValues(t, "10->20->30->40->nil", linkList.Show())
}

func TestAppendElement(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30)
	assert.EqualValues(t, 3, linkList.Size())
	assert.EqualValues(t, "10->20->30->nil", linkList.Show())
}
func TestPrePendElement(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Prepend(5)
	assert.EqualValues(t, 5, linkList.Front())
}

func TestInsertAt(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30)
	err := linkList.InsertAt(1, 40)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 3, linkList.Size())
	assert.EqualValues(t, "10->40->20->30->nil", linkList.Show())
}

func TestInsertAtOutOfRange(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30)
	err := linkList.InsertAt(4, 40)
	assert.NotNil(t, err)
}
func TestSearchElementFound(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30)
	assert.True(t, linkList.Search(20))
}

func TestSearchElementNotFound(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(30).Append(40)
	assert.False(t, linkList.Search(20))
}

func TestRemoveAt(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30)
	assert.EqualValues(t, "10->20->30->nil", linkList.Show())
	err := linkList.RemoveAt(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, linkList.Size())
	assert.EqualValues(t, "10->30->nil", linkList.Show())
}

func TestRemoveAtFront(t *testing.T) {
	linkList := brute_implementation_linklist.NewLinkList(10)
	linkList.Append(20).Append(30)
	assert.EqualValues(t, "10->20->30->nil", linkList.Show())
	err := linkList.RemoveAt(0)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, linkList.Size())
	assert.EqualValues(t, "20->30->nil", linkList.Show())
}

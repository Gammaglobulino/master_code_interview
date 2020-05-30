package brute_implementation_linklist

import (
	"strconv"
	"strings"
)

type DLLNode struct {
	value int
	next  *DLLNode
	prev  *DLLNode
}

type DoubleLinkList struct {
	head *DLLNode
	tail *DLLNode
	size int
}

func NewDoubleLinkList(value int) *DoubleLinkList {
	node := &DLLNode{
		value: value,
		prev:  nil,
		next:  nil,
	}
	return &DoubleLinkList{
		head: node,
		tail: node,
		size: 1,
	}
}

func (ll *DoubleLinkList) Size() int {
	return ll.size
}

func (ll *DoubleLinkList) Append(value int) *DoubleLinkList {
	newNode := &DLLNode{
		value: value,
		prev:  ll.tail,
	}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.size++
	return ll
}

func (ll *DoubleLinkList) Show() string {
	outString := strings.Builder{}
	for i := ll.head; i != nil; i = i.next {
		outString.WriteString(strconv.Itoa(i.value))
		outString.WriteString("->")
	}
	outString.WriteString("nil")
	return outString.String()
}

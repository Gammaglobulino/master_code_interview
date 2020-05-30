package brute_implementation_linklist

import (
	"errors"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type LinkList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkList(value int) *LinkList {
	node := &Node{
		value: value,
	}
	return &LinkList{
		head: node,
		tail: node,
		size: 1,
	}
}
func (ll *LinkList) Size() int {
	return ll.size
}

func (ll *LinkList) Show() string {
	outString := strings.Builder{}
	for i := ll.head; i != nil; i = i.next {
		outString.WriteString(strconv.Itoa(i.value))
		outString.WriteString("->")
	}
	outString.WriteString("nil")
	return outString.String()
}

func (ll *LinkList) Append(value int) *LinkList {
	newNode := &Node{
		value: value,
	}
	ll.tail.next = newNode
	ll.tail = newNode
	ll.size++
	return ll
}

func (ll *LinkList) InsertAt(index int, value int) error {
	if index > ll.Size() {
		return errors.New("Index out of range")
	}
	newNode := &Node{
		value: value,
	}
	if index == 0 {
		ll.Prepend(value)
		return nil
	}
	nextNode := ll.head
	for i := 0; i < index-1; i++ {
		nextNode = nextNode.next
	}
	newNode.next = nextNode.next
	nextNode.next = newNode
	return nil
}

func (ll *LinkList) RemoveAt(index int) error {
	if index > ll.Size() {
		return errors.New("Index out of range")
	}
	if index == 0 {
		ll.head = ll.head.next
		ll.size--
		return nil
	}
	nextNode := ll.head
	for i := 0; i < index-1; i++ {
		nextNode = nextNode.next
	}
	nextNode.next = nextNode.next.next
	ll.size--
	return nil
}

func (ll *LinkList) Search(value int) bool {
	for i := ll.head; i != nil; i = i.next {
		if i.value == value {
			return true
		}
	}
	return false
}

func (ll *LinkList) Prepend(value int) {
	newNode := &Node{
		value: value,
		next:  ll.head,
	}
	ll.head = newNode
}
func (ll *LinkList) Front() int {
	return ll.head.value
}

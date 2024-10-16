package linkedlist

import (
	"fmt"
)

type number interface {
	int | float32 | uint
}

type Node[T number] struct {
	value T
	next  *Node[T]
}

type Linkedlist[T number] struct {
	head   *Node[T]
	tail   *Node[T]
	legnth int
}

// insert item at end
func (l *Linkedlist[T]) append(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.legnth++
}

// insert item at start
func (l *Linkedlist[T]) prepend(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.legnth++
}

func (l *Linkedlist[T]) removeFirst() (Node[T], error) {
	if l.head == nil {
		var empty Node[T]
		return empty, fmt.Errorf("empty linkedlist to remove first")
	}
	node := l.head
	l.head = l.head.next
	node.next = nil
	l.legnth--

	// check if removed item was the only item
	if l.head == nil {
		l.tail = nil
	}

	return *node, nil
}

func (l *Linkedlist[T]) removeLast() (Node[T], error) {
	if l.head == nil {
		var empty Node[T]
		return empty, fmt.Errorf("empty linkedlist to remove last")
	}

	temp := l.head
	previous := l.head

	for temp.next != nil {
		previous = temp
		temp = temp.next
	}
	l.tail = previous
	l.tail.next = nil
	l.legnth--

	// check if removed item was the only item
	if l.legnth == 0 {
		l.head = nil
		l.tail = nil
	}

	return *temp, nil
}

// get node by index
func (l Linkedlist[T]) get(index int) (*Node[T], error) {
	if index < 0 || index >= l.legnth {
		var empty Node[T]
		return &empty, fmt.Errorf("invalid index")
	}

	node := l.head
	i := 0
	for i < index {
		node = node.next
		i++
	}
	return node, nil
}

// set new value to node at informed index
func (l *Linkedlist[T]) set(index int, value T) bool {
	node, err := l.get(index)
	if err != nil {
		return false
	}
	node.value = value
	return true
}

// insert new value to informed index
func (l *Linkedlist[T]) insertAt(index int, value T) bool {
	if index < 0 || index >= l.legnth {
		return false
	}
	//first item
	if index == 0 {
		l.prepend(value)
		return true
	}
	if index == l.legnth {
		l.append(value)
		return true
	}

	node := Node[T]{
		value: value,
		next:  nil,
	}
	temp, _ := l.get(index - 1)
	node.next = temp.next
	temp.next = &node
	l.legnth++

	return true
}

func (l *Linkedlist[T]) removeAt(index int) (Node[T], error) {
	if index < 0 || index >= l.legnth {
		var empty Node[T]
		return empty, fmt.Errorf("can't removeAt when empty")
	}
	if index == 0 {
		node, _ := l.removeFirst()
		return node, nil
	}
	if index == l.legnth {
		node, _ := l.removeLast()
		return node, nil
	}

	previous, _ := l.get(index - 1)
	temp := previous.next
	previous.next = temp.next
	temp.next = nil
	l.legnth--

	return *temp, nil
}

func NewLinkedList[T number]() Linkedlist[T] {
	return Linkedlist[T]{
		head:   nil,
		tail:   nil,
		legnth: 0,
	}
}

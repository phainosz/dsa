package main

import (
	"fmt"
)

type number interface {
	int | int32 | float32 | float64 | uint | uint8
}

type Node[T number] struct {
	value T
	next  *Node[T]
}
type Stack[T number] struct {
	top    *Node[T]
	heigth uint
}

func (s Stack[T]) Peek() {
	if s.heigth == 0 || s.top == nil {
		fmt.Println("Empty stack")
		return
	}
	fmt.Printf("Peeked value: %v\n", s.top.value)
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}

	if s.top != nil {
		node.next = s.top
	}

	s.top = node
	s.heigth++
}

func (s *Stack[T]) Pop() Node[T] {
	if s.top == nil || s.heigth == 0 {
		var nilNode Node[T]
		return nilNode
	}

	node := s.top
	s.top = s.top.next
	node.next = nil
	s.heigth--

	return *node
}

func NewStack[T number]() Stack[T] {
	return Stack[T]{
		heigth: 1,
		top:    nil,
	}
}

func main() {
	stack := NewStack[int]()

	stack.Peek()
	stack.Push(2)
	stack.Push(4)
	stack.Peek()
	node := stack.Pop()
	fmt.Println("Popped item", node.value)
	stack.Peek()
	node = stack.Pop()
	fmt.Println("Popped item", node.value)
	stack.Peek()
}

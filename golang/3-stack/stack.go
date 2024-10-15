package stack

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

func (s Stack[T]) Peek() (T, error) {
	if s.heigth == 0 || s.top == nil {
		fmt.Println("Empty stack")
		return 0, fmt.Errorf("empty stack")
	}
	return s.top.value, nil
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

func (s *Stack[T]) Pop() (Node[T], error) {
	if s.top == nil || s.heigth == 0 {
		var nilNode Node[T]
		return nilNode, fmt.Errorf("stack has no value to pop")
	}

	node := s.top
	s.top = s.top.next
	node.next = nil
	s.heigth--

	return *node, nil
}

func NewStack[T number]() Stack[T] {
	return Stack[T]{
		heigth: 0,
		top:    nil,
	}
}

package queue

import "fmt"

type number interface {
	int | uint | float32
}

type Node[T number] struct {
	next  *Node[T]
	value T
}

type Queue[T number] struct {
	first  *Node[T]
	last   *Node[T]
	length int
}

func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{
		value: value,
		next:  nil,
	}
	if q.first == nil {
		q.first = node
	} else {
		q.last.next = node
	}
	q.last = node
	q.length++
}

func (q *Queue[T]) Dequeue() (Node[T], error) {
	if q.first == nil {
		var empty Node[T]
		return empty, fmt.Errorf("queue has not value do dequeue")
	}
	temp := q.first
	//only one item in the queue
	if q.first == q.last {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.next
		temp.next = nil
	}

	q.length--
	return *temp, nil
}

func (q Queue[T]) Peek() (T, error) {
	if q.first == nil {
		return 0, fmt.Errorf("empty queue")
	}
	return q.first.value, nil
}

func NewQueue[T number]() Queue[T] {
	return Queue[T]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

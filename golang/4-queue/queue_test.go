package queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("testing empty queue", func(t *testing.T) {
		queue := NewQueue[int]()

		if queue.length != 0 {
			t.Fatal("queue should've been empty")
		}

		_, err := queue.Peek()
		if err == nil {
			t.Error("Expected an error when peek an empty queue")
		}

		_, err = queue.Dequeue()
		if err == nil {
			t.Error("Expected an error when dequeue empty queue")
		}

		if queue.first != nil {
			t.Error("Expected nil first element in empty queue")
		}

		if queue.last != nil {
			t.Error("Expected nil first element in empty queue")
		}
	})

	t.Run("testing int queue", func(t *testing.T) {
		queue := NewQueue[int]()

		queue.Enqueue(1)
		queue.Enqueue(4)
		queue.Enqueue(7)
		queue.Enqueue(70)

		value, err := queue.Peek()
		validateValue(t, value, 1, err)

		node, err := queue.Dequeue()
		if err != nil {
			t.Error("did not expected error when dequeue")
		}
		if node.value != 1 {
			t.Errorf("Expect 1 but got %d", node.value)
		}

		node, err = queue.Dequeue()
		validateValue(t, node.value, 4, err)

		value, err = queue.Peek()
		validateValue(t, value, 7, err)

		node, err = queue.Dequeue()
		validateValue(t, node.value, 7, err)

		value, err = queue.Peek()
		validateValue(t, value, 70, err)
	})
}

func validateValue(t testing.TB, value, expectedValue int, err error) {
	t.Helper()
	if err != nil {
		t.Error("did not expected error")
	}
	if value != expectedValue {
		t.Errorf("Expect %d but got %d", expectedValue, value)
	}
}

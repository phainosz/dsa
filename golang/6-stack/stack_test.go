package stack

import "testing"

func TestStack(t *testing.T) {
	t.Run("testing empty stack", func(t *testing.T) {
		stack := NewStack[int]()

		if stack.heigth != 0 {
			t.Errorf("Expect 0 but got %d", stack.heigth)
		}

		if stack.top != nil {
			t.Errorf("Expect nil but got %v", stack.heigth)
		}

		value, err := stack.Peek()
		if err == nil {
			t.Errorf("expected an error but got %d", value)
		}

		_, err = stack.Pop()
		if err == nil {
			t.Fatal("expected error when poping empty stack")
		}
	})

	t.Run("testing int stack", func(t *testing.T) {
		stack := NewStack[int]()

		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if stack.heigth != 3 {
			t.Errorf("expected stack heigth equal 3 but got %d", stack.heigth)
		}

		node, err := stack.Pop()
		if err != nil {
			t.Error("did not expect an error when poping from stack")
		}

		if node.value != 3 {
			t.Errorf("expect %d but got %d", 3, node.value)
		}

		value, err := stack.Peek()
		if err != nil {
			t.Error("did not expected an error when peeking from stack")
		}

		if value != 2 {
			t.Errorf("expect %d but got %d", 3, node.value)
		}
	})

}

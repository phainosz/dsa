package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {

	t.Run("testing empty linkedlist", func(t *testing.T) { emptyTest(t) })
	t.Run("testing append to linkedlist", func(t *testing.T) { appendTest(t) })
	t.Run("testing prepend to linkedlist", func(t *testing.T) { prependTest(t) })
	t.Run("testing removeFirst to linkedlist", func(t *testing.T) { removeFirstTest(t) })
	t.Run("testing removeLast to linkedlist", func(t *testing.T) { removeLastTest(t) })
	t.Run("testing insertAt to linkedlist", func(t *testing.T) { insetAtTest(t) })
}

func emptyTest(t *testing.T) {
	t.Helper()

	ll := NewLinkedList[int]()
	testedValue := 0
	testedIndex := 1

	_, err := ll.get(testedIndex)
	if err == nil {
		t.Error("expected error when get in an empty linkedlist")
	}

	testedIndex = 1
	_, err = ll.removeAt(testedIndex)
	if err == nil {
		t.Error("expected error when removeAt in an empty linkedlist")
	}

	_, err = ll.removeFirst()
	if err == nil {
		t.Error("expected error when removeFirst in an empty linkedlist")
	}

	_, err = ll.removeLast()
	if err == nil {
		t.Error("expected error when removeLast in an empty linkedlist")
	}

	testedValue = 7
	testedIndex = 10
	result := ll.set(testedIndex, testedValue)
	if result {
		t.Error("expected false when set in an empty linkedlist")
	}
}

func appendTest(t *testing.T) {
	t.Helper()

	ll := NewLinkedList[int]()
	testedValue := 1
	testedIndex := 0

	ll.append(testedValue)
	node, err := ll.get(testedIndex)
	if err != nil {
		t.Error("did not expected error when get")
	}
	if node.value != testedValue {
		t.Errorf("expected %d got %d", testedValue, node.value)
	}

	testedValue = 2
	testedIndex = 0
	ll.append(testedValue)
	testedValue = 10
	result := ll.set(testedIndex, testedValue)
	if !result {
		t.Errorf("expect set to return true but got %v", result)
	}
	node, err = ll.get(testedIndex)
	if err != nil {
		t.Error("did not expected error when get")
	}
	if node.value != testedValue {
		t.Errorf("expected %d got %d", testedValue, node.value)
	}
}

func prependTest(t *testing.T) {
	t.Helper()

	ll := NewLinkedList[int]()
	testedValue := 0
	testedIndex := 0

	ll.prepend(testedValue)
	node, err := ll.get(testedIndex)
	if err != nil {
		t.Error("did not expected error when get")
	}
	if node.value != testedValue {
		t.Errorf("expected %d got %d", testedValue, node.value)
	}

	testedValue = 2
	testedIndex = 0
	ll.prepend(testedValue)
	node, err = ll.get(testedIndex)
	if err != nil {
		t.Error("did not expected error when get")
	}
	if node.value != testedValue {
		t.Errorf("expected %d got %d", testedValue, node.value)
	}

	testedValue = 0
	testedIndex = 1
	node, err = ll.get(testedIndex)
	if err != nil {
		t.Error("did not expected error when get")
	}
	if node.value != testedValue {
		t.Errorf("expected %d got %d", testedValue, node.value)
	}
}

func removeFirstTest(t *testing.T) {
	t.Helper()

	ll := NewLinkedList[int]()

	testsData := []int{1, 2, 3}

	for _, testData := range testsData {
		ll.append(testData)
	}

	for _, testData := range testsData {
		node, err := ll.removeFirst()
		if err != nil {
			t.Error("did not expected error when removeFirst")
		}
		if node.value != testData {
			t.Errorf("expected %d got %d", testData, node.value)
		}
	}
}

func removeLastTest(t *testing.T) {
	t.Helper()

	ll := NewLinkedList[int]()

	testsData := []int{1, 2, 3}

	for _, testData := range testsData {
		ll.prepend(testData)
	}

	for _, testData := range testsData {
		node, err := ll.removeLast()
		if err != nil {
			t.Error("did not expected error when removeLast")
		}
		if node.value != testData {
			t.Errorf("expected %d got %d", testData, node.value)
		}
	}
}

func insetAtTest(t *testing.T) {
	t.Helper()

	ll := NewLinkedList[int]()

	ll.append(1)
	ll.append(2)
	ll.append(3)

	result := ll.insertAt(0, 10)
	if !result {
		t.Errorf("expect insertAt to return true but got %v", result)
	}

	node, err := ll.get(0)
	if err != nil {
		t.Error("did not expected error when get")
	}
	if node.value != 10 {
		t.Errorf("expected %d got %d", 10, node.value)
	}

}

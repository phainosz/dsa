package recursion

import "testing"

type testData struct {
	value    int
	expected int
}

func TestFibonacci(t *testing.T) {
	t.Run("testing fibonacci", func(t *testing.T) {
		tests := []testData{
			{value: 0, expected: 0},
			{value: 1, expected: 1},
			{value: 2, expected: 1},
			{value: 3, expected: 2},
			{value: 4, expected: 3},
			{value: 5, expected: 5},
			{value: 9, expected: 34},
			{value: 11, expected: 89},
		}

		for _, testData := range tests {
			result := fibonacci(testData.value)
			if testData.expected != result {
				t.Errorf("got %d wanted %d", result, testData.expected)
			}
		}

	})
}

func TestFactorial(t *testing.T) {
	t.Run("testing factorial", func(t *testing.T) {
		tests := []testData{
			{value: 0, expected: 1},
			{value: 1, expected: 1},
			{value: 2, expected: 2},
			{value: 3, expected: 6},
			{value: 4, expected: 24},
			{value: 5, expected: 120},
			{value: 9, expected: 362880},
			{value: 11, expected: 39916800},
		}

		for _, testData := range tests {
			result := factorial(testData.value)
			if testData.expected != result {
				t.Errorf("got %d wanted %d", result, testData.expected)
			}
		}

	})
}

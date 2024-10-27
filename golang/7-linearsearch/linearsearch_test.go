package linearsearch

import "testing"

type dataTest struct {
	needle        int
	expectedIndex int
	expectedError bool
}

var collection = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestLinearSearch(t *testing.T) {
	tests := []dataTest{
		{needle: 0, expectedIndex: 0, expectedError: false},
		{needle: 1, expectedIndex: 1, expectedError: false},
		{needle: 2, expectedIndex: 2, expectedError: false},
		{needle: 10, expectedIndex: -1, expectedError: true},
		{needle: 11, expectedIndex: -1, expectedError: true},
	}

	for _, test := range tests {
		indexFound, err := linearsearch(collection, test.needle)

		if test.expectedError && err == nil {
			t.Errorf("expected error for needle %d", test.needle)
		}

		if test.expectedIndex != indexFound {
			t.Errorf("expected index %d but found %d", test.expectedIndex, indexFound)
		}

	}

}

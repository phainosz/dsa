package linearsearch

import "testing"

type dataTest struct {
	needle        int
	expectedIndex int
	expectedError bool
}

var collection = []int{9, 19, 10, 14, 15, 0, 2, 3}

func TestLinearSearch(t *testing.T) {
	tests := prepareDataTest(collection)

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

func prepareDataTest(arr []int) []dataTest {
	var tests []dataTest
	for i, v := range arr {
		test := dataTest{v, i, false}
		tests = append(tests, test)
	}
	tests = append(tests, dataTest{11, -1, true})
	tests = append(tests, dataTest{99, -1, true})
	tests = append(tests, dataTest{999, -1, true})

	return tests
}

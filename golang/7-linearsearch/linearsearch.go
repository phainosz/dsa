package linearsearch

import (
	"fmt"
)

type number interface {
	int | uint | float32
}

func linearsearch[T number](array []T, needle T) (int, error) {
	for index, value := range array {
		if needle == value {
			return index, nil
		}
	}
	return -1, fmt.Errorf("item %v not found", needle)
}

package binarysearch

import "fmt"

func binarySearch[T int](array []T, needle T) (int, error) {
	low := 0               //low or left
	high := len(array) - 1 //high ir right

	for low <= high {
		middle := low + (high-low)/2

		//check if needle is the middle element
		if array[middle] == needle {
			return middle, nil
		}

		//if needle is bigger, ignore left side
		if array[middle] < needle {
			low = middle + 1

			//else, ignore right side
		} else {
			high = middle - 1
		}
	}

	//not found if reached here
	return -1, fmt.Errorf("item %v not found", needle)
}

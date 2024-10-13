package main

import "fmt"

func main() {
	fmt.Println("Array with fized size")
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	for index, element := range array {
		fmt.Printf("Index: %d element: %d\n", index, element)
	}

	fmt.Println("Array with dynamic size or as called in Go, slice")
	var slice []int = []int{1, 2, 3, 4, 5}

	for index, element := range slice {
		fmt.Printf("Index: %d element: %d\n", index, element)
	}
}

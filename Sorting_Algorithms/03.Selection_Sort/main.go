package main

import "fmt"

func main() {
	// Example slice to be sorted
	slice := []int{64, 34, 25, 22, 11, 90}

	fmt.Println("Unsorted slice:", slice)

	selectionSort(slice)

	fmt.Println("Sorted slice:", slice)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find the minimum element in the unsorted portion
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

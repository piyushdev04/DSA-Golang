package main

import "fmt"

// flip reverses the order of first k elemnts in array
func flip(arr []int, k int) {
	for i := 0; i < k/2; i++ {
		arr[i], arr[k-i-1] = arr[k-i-1], arr[i]
	}
}

// findMax returns the index of the maximum element in arr[0..n-1]
func findMax(arr []int, n int) int {
	maxIdx := 0
	for i := 1; i < n; i++ {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func pancakeSort(arr []int) {
	for currSize := len(arr); currSize > 1; currSize-- {
		// Find the index of the maximum element in arr[0..currSize-1]
		maxIdx := findMax(arr, currSize)

		// Move the maximum element to the end of the current array
		if maxIdx != currSize-1 {
			// Flip the maximum element to the front
			flip(arr, maxIdx+1)
			// Flip it to the end of the current array
			flip(arr, currSize)
		}
	}
}

func main() {
	slice := []int{3, 6, 2, 7, 5, 8, 4}

	fmt.Println("Unsorted slice:", slice)

	pancakeSort(slice)

	fmt.Println("Sorted slice:", slice)
}

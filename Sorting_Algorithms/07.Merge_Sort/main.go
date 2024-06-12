package main

import "fmt"

func main() {
	// Example slice to be sorted
	slice := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted slice:", slice)

	sortedSlice := mergeSort(slice)

	fmt.Println("Sorted slice:", sortedSlice)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr)/2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// merge merges two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i,j := 0,0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left and right slices
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
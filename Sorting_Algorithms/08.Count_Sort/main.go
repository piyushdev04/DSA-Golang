package main

import "fmt"

func countingSort(arr []int) []int {
	// Find the maximum element in the array
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Create a count array to store the count of each unique object
	count := make([]int, max+1)

	// Store the count of each element
	for _, num := range arr {
		count[num]++
	}

	// Store the sorted elements in the result array
	result := make([]int, len(arr))
	index := 0
	for i, cnt := range count {
		for cnt > 0 {
			result[index] = i
			index++
			cnt--
		}
	}

	return result
}

func main() {
	slice := []int{4, 2, 2, 8, 3, 3, 1}

	fmt.Println("Unsorted slice:", slice)

	sortedSlice := countingSort(slice)

	fmt.Println("Sorted slice:", sortedSlice)
}

package main

import "fmt"

func main() {
	// Example slice to be sorted
	slice := []int{64, 34, 25, 12, 11, 90}

	fmt.Println("Unsorted slice:", slice)

	insertionSort(slice)

	fmt.Println("Sorted slice:", slice)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		/*Move elements of arr[0..i-1], that are greater than key,
		to one position ahead of their current position*/
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

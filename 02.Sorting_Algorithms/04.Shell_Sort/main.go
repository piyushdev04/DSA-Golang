package main

import "fmt"

func main() {
	// Example slice to br sorted
	slice := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted slice:", slice)

	shellSort(slice)

	fmt.Println("Sorted slice:", slice)
}

func shellSort(arr []int) {
	n := len(arr)
	// start with a big gap, then reduce the gap
	for gap := n/2; gap > 0; gap /= 2 {
		// Perform a gapped insertion sort
		for i := gap; i < n; i++ {
			// Save arr[i] to temporary variable and make a hole at position i
			temp := arr[i]
			j := i
			// shift earlier gap-sorted elements up until the corrected location for arr[i] is found
			for ; j >= gap && arr[j-gap] > temp; j -= gap{
				arr[j] = arr[j-gap]
			}
			// Put temp (the originalarr[i]) in its correct location
			arr[j] = temp
		}
	}
}
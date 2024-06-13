package main

import "fmt"

func main() {
	// Example slice to be sorted
	slice := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted slice:", slice)

	combSort(slice)

	fmt.Println("Sorted slice:", slice)
}

func combSort(arr []int) {
	n := len(arr)
	gap := n
	shrink := 1.3
	sorted := false

	for !sorted {
		// Update the gap value for a next comb
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		}

		// Compare all elements with current gap
		for i := 0; i+gap < n; i++ {
			if  arr[i] > arr[i+gap] {
				// Swap arr[i] and arr[i+gap]
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				sorted = false
			}
		}
	}
}
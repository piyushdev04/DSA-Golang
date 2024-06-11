package main

import (
	"math/rand"
	"fmt"
)

func main() {
	// Example slice to be sorted
	slice := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted slice:", slice)

	quicksort(slice)

	fmt.Println("Sorted slice:", slice)
}

func quicksort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	//Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the end
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Partition
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	// Move the pivot to its final place
	a[left], a[right] = a[right], a[left]

	// Recursively apply to the sub-arrays
	quicksort(a[:left])
	quicksort(a[left+1:])
}

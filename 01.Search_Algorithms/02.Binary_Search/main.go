package main

import "fmt"

func binarySearch(target int, array []int) bool {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] < target {
			low = mid + 1
		} else if array[mid] > target {
			high = mid - 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 30, 40, 100}
	fmt.Println(binarySearch(30, array))
	fmt.Println(binarySearch(10, array))
}
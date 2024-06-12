package main

import "fmt"

// A utility function tom get the maximum value in an array
func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// A function to do counting sort of arr[] according toth edigit represented by exp.
func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n) // output array
	count := make([]int, 10) // count array to store the count of occurrences of digits (0-9)

	// Store count of occurrences of each digit in the count array
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	// Change count[i] so that it contains the actual position of this digit in the output array
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	//copy the output array to arr[] now contains sorted numbers according to the current digit
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func radixSort(arr []int) {
	// Find the maximum number to know the number of digits
	m := getMax(arr)

	// Do counting sort for every digit. Note that instead of passing the digit number, exp is passed
	// exp is 10^i where i is the current digit number
	for exp := 1; m/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}

func main() {
	slice := []int{170, 45, 75, 90, 802, 24, 2, 66}

	fmt.Println("Unsorted slice:", slice)

	radixSort(slice)

	fmt.Println("Sorted slice:", slice)
}

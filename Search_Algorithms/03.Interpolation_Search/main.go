package main

import "fmt"

func interpolationSearch(array []int, key int) int {

	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if key < min {
			return low
		}

		if key > max {
			return high + 1
		}

		// make a guess of the location
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
			guess = low + offset
		}

		// maybe we found it?
		if array[guess] == key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}

func main() {
	items := []int{1,9,3,6,4,8,2,5}
	fmt.Println(interpolationSearch(items, 5))
}

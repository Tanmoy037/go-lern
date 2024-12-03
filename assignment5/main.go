// Challenge 1: Sort a Slice
// Take a slice of integers as input.
// Sort the slice using Bubble Sort.
// Print the sorted slice.
// Challenge 2: Search for an Element
// Use the sorted slice from Challenge 1.
// Implement Binary Search to find the position of a given number.
// Print the position if the number is found, or a message saying it’s not found.
// Challenge 3: Find the Median
// Sort the slice.
// Find the median of the slice:
// For odd-sized slices, the median is the middle element.
// For even-sized slices, the median is the average of the two middle elements.
package main

import "fmt"


func bubbleShort(slice []int) [] int {
	length := len(slice)

	for i := 0; i < length ; i++ {
		
		for j := 0; j < length -i-1; j++ {
		
			if slice[j] > slice[j+1] {

				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
} 

func binarySearch(slice [] int, target int) int {
	left , right := 0, len(slice) - 1

	for left <= right {
		mid := left + (right - left)/2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func medianOfSlice(slice [] int) float64 {
	length := len(slice)
	if length%2 == 1 {
		return float64(slice[length/2])
	} else {
		mid1, mid2 := slice[length/2-1] , slice[length/2]
		return float64(mid1+mid2) / 2.0
	}

}


func main (){
	slice := []int {5, 3, 8, 4, 35 }
	result := bubbleShort(slice)
	fmt.Println(result)
	target := 8
	respose := binarySearch(result, target)
	fmt.Println(respose)
	median := medianOfSlice(slice)
	fmt.Println(median)
}
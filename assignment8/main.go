// Challenge 1: Sort an Array
// Use merge sort to sort an array of integers.
// Input: [38, 27, 43, 3, 9, 82, 10]
// Output: [3, 9, 10, 27, 38, 43, 82]
// Challenge 2: Find Kth Smallest Element
// After sorting the array, return the 
// 𝑘
// k-th smallest element.
// Input: [3, 9, 10, 27, 38, 43, 82], 
// 𝑘
// =
// 4
// k=4
// Output: 27
package main

import "fmt"
func mergeSort(arr [] int ) [] int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left , right [] int ) [] int {
	result := [] int {}
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
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func findKthElement(arr [] int, k int ) int {
	response := mergeSort(arr)
	return response[k-1]
}

func main () {
	arr := [] int {38, 27, 43, 3, 9, 82, 10}
	result := mergeSort(arr)
	fmt.Println(result)

	k := 4
	kthElement := findKthElement(result, k)
	fmt.Println(kthElement)
}
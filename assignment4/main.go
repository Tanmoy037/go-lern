// Today's Challenges
// Challenge 1: Sum of Elements
// Write a program that:

// Accepts a slice of integers.
// Calculates and returns the sum of all elements.

// Challenge 2: Find the Second Largest Number
// Accept a slice of integers.
// Return the second largest number in the slice.

// Challenge 3: Reverse a Slice
// Accept a slice of integers.
// Return a new slice with the elements reversed.

package main

import (
	"fmt"
)

func sumOfElements(slice []int) int {
	
	sum := 0

	for _, val := range slice {
		sum += val
	}
	return sum

}

func secondLargestInteger(slice []int) int {
	largestElement := -1 << 31
	secondLargestElement := -1 << 31
	for _, number := range slice {
		if number > largestElement {
			secondLargestElement = largestElement
			largestElement = number
		} else if number > secondLargestElement && number != largestElement {
			secondLargestElement = number
		}
	}
	return secondLargestElement
}

func reverseSlice(slice []int) []int {
	length := len(slice)
	reseved := make([]int, length)

	for i := 0; i < length; i++ {
		reseved[i] = slice[length-1-i]
	}
	return reseved
}

func main (){

	var length int 

	fmt.Print("Enter the lenth of the slice : ")
	fmt.Scanln(&length)

	slice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanln(&slice[i])
	}
	sum := sumOfElements(slice)
	fmt.Printf("The sum of the slice elements is: %d\n", sum)



	secondLargest := secondLargestInteger(slice)
	if secondLargest != -1 {
		fmt.Printf("The second largest number is: %d\n", secondLargest)
	} else {
		fmt.Println("No second largest number found.")
	}

	reversed := reverseSlice(slice)
	fmt.Println("Reversed Slice:", reversed)

}
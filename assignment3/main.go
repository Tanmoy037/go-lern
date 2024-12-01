// Today's Assignment
// Write a program with a function checkEvenOdd (as above), but modify it to:

// Return a string ("Even" or "Odd") instead of printing directly.
// Write another program:

// Create an array of 5 numbers.
// Find the largest number in the array using a loop.


package main

import "fmt"
func checkEvenOdd(number int) string {
	if number % 2 == 0{
		return "Even"
	} else {
		return "Odd"
	}
}

func checkLangestNumber() int {

	var arr [5] int = [5] int {10, 20 , 30 , 40 , 50}
	largestElement := arr[0]
	for  _, value := range arr {
		
		if largestElement < value {
			largestElement = value
		}
	}
	return largestElement
}

func main (){
	fmt.Println(checkEvenOdd(10))
	fmt.Println(checkLangestNumber())
}
// Today's Assignment
// Write a program that:

// Accepts a number as input.
// Prints "Even" if the number is even, otherwise "Odd".
// Write another program to:

// Print numbers from 1 to 10.
// Skip printing 5 (use the continue statement).

package main

import "fmt"

func main(){
	var number int

	fmt.Print("Enter the number : ")
	fmt.Scanln(&number)

	if number % 2 == 0{
		fmt.Println("This is a Even number")
	} else {
		fmt.Println("This is a Odd number")
	}
	

	for i := 1; i <= 10; i++ {
		if i==5 {
			continue;
		} else {
			fmt.Println(i)
		}
	}

}
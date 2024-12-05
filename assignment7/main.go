// Challenge 1: Sum of Digits
// Write a recursive function to calculate the sum of digits of a number.
// Input: 1234
// Output: 10 (1 + 2 + 3 + 4)
// Challenge 2: Power of a Number
// Write a recursive function to calculate 
// 𝑎
// 𝑏
// a 
// b
//   (a raised to the power b).
// Input: a = 2, b = 3
// Output: 8 (2^3)
// Challenge 3: Greatest Common Divisor (GCD)
// Write a recursive function to find the GCD of two numbers using Euclid’s algorithm.
// Input: a = 48, b = 18
// Output: 6

package main

import "fmt"


func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n % 10 + sumOfDigits(n/10)
	
}

func powerOfANumber(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * powerOfANumber(a, b-1)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main () {
	num := 1234
	fmt.Printf("Sum of digits of %d is: %d\n", num, sumOfDigits(num))

	base, exponent := 2,3 
	fmt.Printf("%d^%d is: %d\n", base, exponent, powerOfANumber(base, exponent))

	a, b := 48, 18
    fmt.Printf("GCD of %d and %d is: %d\n", a, b, gcd(a, b))
}
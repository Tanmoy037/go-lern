// Challenge 1: Word Frequency Counter
// Accept a paragraph of text as input.
// Count the occurrences of each word using a map.
// Print each word and its count.
// Challenge 2: Reverse a String
// Accept a string as input.
// Return the string reversed.
// Input: "Hello"
// Output: "olleH"
// Challenge 3: Palindrome Checker
// Accept a string as input.
// Check if the string is a palindrome (ignoring spaces and capitalization).
// Input: "madam"
// Output: true
package main

import (
	"fmt"
	"strings"
)

func frequencyCounter (paragraph string) map[string]int {
	words := strings.Fields(paragraph)
	frequency := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		frequency[word]++
	}
	return frequency
}

func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2 ; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i] , runes[i]
	}

	return string(runes)

}

func isPalindrome(str string) bool {
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	reversed := reverseString(str)
	return str==reversed
}



func main () {
	paragraph := "I like software engineering"
	frequency := frequencyCounter(paragraph)
	fmt.Println(frequency)

	str := "Hello"
	fmt.Println("Reversed String: ", reverseString(str))

	palindrome := "Madam"
	fmt.Println("Is Palindrome", isPalindrome(palindrome))

}
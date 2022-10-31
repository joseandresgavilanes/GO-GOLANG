package main

import "fmt"
import "strings"

func isPalindrome(text string) {
	var textReverse string

	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("It's a palindrome")
	} else {
		fmt.Println("It isn't a palindrome")
	}
}



func main(){
	isPalindrome("Ana")
	isPalindrome("Patrick")
	isPalindrome("NaN")
	
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palindromeCheck(str string) bool {
	str = strings.TrimSpace(str)

	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {

	fmt.Print("Enter a string: ")

	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	palindrome := palindromeCheck(str)

	if palindrome {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func WordFrequencyCount(text string) map[string]int {
	wordCount := make(map[string]int)

	text = strings.ToLower(text)

	var cleaned strings.Builder

	for _, ch := range text {
		if unicode.IsLetter(ch) || unicode.IsSpace(ch) {
			cleaned.WriteRune(ch)
		} else {
			cleaned.WriteRune(' ')
		}
	}

	text = cleaned.String()

	for _, word := range strings.Fields(text) {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	fmt.Print("Enter a string: ")

	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	wordCount := WordFrequencyCount(str)

	fmt.Println("Word Frequencies:")

	for word, count := range wordCount {
		fmt.Printf("- %s: %d\n", word, count)
	}
}

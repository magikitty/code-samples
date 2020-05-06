package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	wordCount("The happy meerkat and the happy mongoose play together! They are such a happy meerkat and mongoose. So happy, happy, yes, indeed!")
	wordCount("Seven little sheep jumped over seven little hills. They jumped high! As high as a little sheep could jump! High over the hills...")
}

// Returns map with all words in content and number of times each word is in content
func wordCount(content string) map[string]int {
	contentLower := strings.ToLower(content)
	contentStripped := stripPunctuation(contentLower)
	words := strings.Fields(contentStripped)
	wordMap := make(map[string]int)

	for _, word := range words {
		_, ok := wordMap[word]
		if ok == true {
			oldValue := wordMap[word]
			newValue := oldValue + 1
			wordMap[word] = newValue
		} else {
			wordMap[word] = 1
		}
	}
	fmt.Println(wordMap)
	return wordMap
}

func stripPunctuation(content string) string {
	runeContent := []rune(content)
	var strippedContent string

	for _, char := range runeContent {
		if unicode.IsPunct(char) != true {
			strippedContent += string(char)
		}
	}
	return strippedContent
}

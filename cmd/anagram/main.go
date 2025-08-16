package main

import (
	"fmt"
	"github.com/aliskhannn/anagram/internal/anagram"
)

func main() {
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	anagrams := anagram.FindAnagrams(dictionary)

	for key, words := range anagrams {
		fmt.Println("Anagram group for:", key)
		for _, word := range words {
			fmt.Println(" -", word)
		}
	}
}

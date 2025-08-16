package anagram

import (
	"sort"
	"strings"
)

// FindAnagrams finds all sets of anagrams in the given dictionary.
// It returns a map where the key is the first encountered word of each anagram set,
// and the value is a sorted slice of all words that are anagrams of each other.
// Words with no anagrams are excluded from the result.
//
// The input dictionary is normalized to lowercase before processing.
// The function ensures that duplicate words in the dictionary are ignored.
func FindAnagrams(dictionary []string) map[string][]string {
	if len(dictionary) == 0 {
		return map[string][]string{} // return an empty map if the dictionary is empty
	}

	// A map to hold groups of words that are anagrams of each other.
	// The key is the sorted character representation of the word,
	// and the value is a set of words that match that sorted representation.
	groups := make(map[string]map[string]struct{})

	for _, word := range dictionary {
		word = strings.ToLower(word)   // normalize word to lowercase
		sortedWord := sortString(word) // sort the characters in the word

		if _, exists := groups[sortedWord]; !exists {
			groups[sortedWord] = make(map[string]struct{}) // create a new set for this sorted word
		}
		groups[sortedWord][word] = struct{}{} // add the word to the set
	}

	anagrams := make(map[string][]string)

	for _, wordsSet := range groups {
		// Convert the set of words to a slice
		words := make([]string, 0, len(wordsSet))

		for word := range wordsSet {
			words = append(words, word)
		}

		// If there are multiple words that are anagrams, sort them and add to the result.
		if len(words) > 1 {
			sort.Strings(words)
			anagrams[words[0]] = words // use the first word as the key
		}
	}

	return anagrams
}

// sortString returns a string with its characters sorted in ascending order.
func sortString(word string) string {
	if len(word) < 2 {
		return word // no need to sort if the word is empty or has one character
	}

	// Convert the string to a slice of runes.
	runes := []rune(word)

	// Sort the runes in ascending order.
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convert the sorted runes back to a string.
	return string(runes)
}

package isogram

import "strings"

func IsIsogram(word string) bool {

	var lowerWord = strings.ToLower(word)
	var ignoreCharacters = "- "
	// Create a map to store the frequency of each character
	charFreq := make(map[rune]int)

	// Iterate over each character in the string
	for _, char := range lowerWord {
		// If the character is already in the map, it's not an isogram
		if charFreq[char] > 0 && !strings.ContainsRune(ignoreCharacters, char) {
			return false
		}

		// Increment the frequency of the character in the map
		charFreq[char]++
	}

	// If we've made it this far, the string is an isogram
	return true
}

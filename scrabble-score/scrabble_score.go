package scrabble

import "strings"

func Score(word string) int {
	var count = 0
	for _, c := range word {
		if strings.ContainsRune("AEIOULNRSTaeioulnrst", c) {
			count = count + 1
		}
		if strings.ContainsRune("DGdg", c) {
			count = count + 2
		}
		if strings.ContainsRune("BCMPbcmp", c) {
			count = count + 3
		}
		if strings.ContainsRune("FHVWYfhvwy", c) {
			count = count + 4
		}
		if strings.ContainsRune("Kk", c) {
			count = count + 5
		}
		if strings.ContainsRune("JXjx", c) {
			count = count + 8
		}
		if strings.ContainsRune("QZqz", c) {
			count = count + 10
		}
	}
	return count
}

// ["A", "E", "I", "O", "U", "L", "N", "R", "S", "T"]

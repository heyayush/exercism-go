package luhn

import "unicode"

func Valid(id string) bool {
	// filter only the digits
	for _, char := range id {
		// If the character is not a digit, return false
		if !unicode.IsDigit(char) {
			return false
		}
	}
}

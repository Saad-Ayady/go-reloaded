package mycode

import "strings"

func ChangAtoAn(s []string) []string {
	for i, val := range s {
		// check a char
		if val == "a" || val == "A" && i > 0 && i < len(s)-1 {
			// get char next
			nextWord := s[i+1]
			// list of char
			vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
			for _, char := range vowels {
				// check if first char in nextWord == char like nextWord[0] == char
				if strings.HasPrefix(nextWord, char) {
					s[i] = val + "n"
				}
			}
		}
	}
	return s
}

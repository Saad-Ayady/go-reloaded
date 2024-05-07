package mycode

import (
	"strconv"
	"strings"
)

func UpLowCap_numpers(s []string) []string {
	for idx, word := range s {
		// Check if word = "(up," 
		if word == "(up," && idx < len(s)-1{
			// Extract the value after "(up," and convert it to an integer.
			endWord := strings.TrimSuffix(s[idx+1], ")") // check ) in s
			nb, _ := strconv.Atoi(endWord)
			for i := 1; i <= nb; i++ {
				calcu := (idx - i)
				if calcu >= 0 && calcu < len(s) {
					s[calcu] = strings.ToUpper(s[calcu])
				}
			}
			s = r2Index(s, idx)
		} else if word == "(low," {
			endWord := strings.TrimSuffix(s[idx+1], ")")
			nb, _ := strconv.Atoi(endWord)
			for i := 1; i <= nb; i++ {
				calcu := (idx - i)
				if calcu >= 0 && calcu < len(s) {
					s[calcu] = strings.ToLower(s[calcu])
				}
			}
			s = r2Index(s, idx)
		} else if word == "(cap," {
			endWord := strings.TrimSuffix(s[idx+1], ")")
			nb, _ := strconv.Atoi(endWord)
			for i := 1; i <= nb; i++ {
				calcu := (idx - i)
				if calcu >= 0 && calcu < len(s) {
					s[calcu] = strings.Title(s[calcu])
				}
			}
			s = r2Index(s, idx)
		}
	}
	return s
}

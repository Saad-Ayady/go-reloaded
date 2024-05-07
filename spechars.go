package mycode

import "strings"

func SpichelCarecter(s []string) []string {
	speChars := []string{".", ",", "!", "?", ":", ";"}
	for idx := 1; idx < len(s); idx++ {
		for j := 0; j < len(speChars); j++ {
			if s[idx] == speChars[j] {
				s[idx-1] = s[idx-1] + speChars[j]
				s = append(s[:idx], s[idx+1:]...)
				idx--
			} else if strings.HasPrefix(s[idx], speChars[j]) { // chack if furst char in s == furst char in speChars
				s[idx-1] = s[idx-1] + speChars[j]
				s[idx] = s[idx][1:]
			}
		}
	}
	return s
}

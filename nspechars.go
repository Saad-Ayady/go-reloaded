package mycode

import "strings"

func MultChar(s []string) []string {
	spChars := []string{",", "!", "?", ":", ";", "."}
	for i, val := range s {
		for _, v := range spChars {
			if len(val) > 1 && strings.HasPrefix(val, v) && strings.HasSuffix(val, v) { // chack furst and lasst char 
				if i > 0 {
					s[i-1] = s[i-1] + val
					s = rIndex(s, i)
				} 
			}
		}
	}
	return SpichelCarecter(s)
}

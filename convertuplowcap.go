package mycode

import "strings"

func UpLowCap(s []string) []string {
	for i, tag := range s {
		if tag == "(low)" {
			if i > 0 && i < len(s) {
				s[i-1] = strings.ToLower(s[i-1])
				// remov word 
				s[i] = strings.ReplaceAll(s[i], "(low)", "")
			} else if i == 0 {
				s[0] = strings.ReplaceAll(s[0], "(low)", "")
			} else {
				s[i] = strings.ReplaceAll(s[i], "(low)", "")
			}
		} else if tag == "(up)" {
			if i > 0 && i < len(s) {
				s[i-1] = strings.ToUpper(s[i-1])
				s[i] = strings.ReplaceAll(s[i], "(up)", "")
			} else if i == 0 {
				s[0] = strings.ReplaceAll(s[0], "(up)", "")
			} else {
				s[i] = strings.ReplaceAll(s[i], "(up)", "")
			}

		} else if tag == "(cap)" {
			if i > 0 && i < len(s) {
				s[i-1] = strings.Title(s[i-1])
				s[i] = strings.ReplaceAll(s[i], "(cap)", "")
			} else if i == 0 {
				s[0] = strings.ReplaceAll(s[0], "(cap)", "")
			} else {
				s[i] = strings.ReplaceAll(s[i], "(cap)", "")
			}
		}
	}
	return s
}

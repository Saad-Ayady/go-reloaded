package mycode

import "strings"

func Split(str string) []string {
	var Spr []string
	var word string

	for _, ch := range str {
		// check TAP and SPIS
		if ch == ' ' || ch == '\t' {
			if word != "" {
				Spr = append(Spr, word)
				word = ""
			}
			// check new line
		} else if ch == '\n' {
			if word != "" {
				Spr = append(Spr, word)
				word = ""
			}
			Spr = append(Spr, "\n") // add new line to array
		} else {
			word += string(ch)
		}
	}
	// check if a nermal char 
	if word != "" {
		Spr = append(Spr, word)
	}
	return Spr
}

func Join(str []string) string {
	var result strings.Builder
	for i, field := range str {
		// 
		if i == 0 || field == "\n" || str[i-1] == "\n" {
			result.WriteString(field) // Write field in result and not add spice
		} else {
			result.WriteString(" " + field) // Write field in result and add spice
		}
	}
	return result.String()
}

package mycode

func rIndex(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func r2Index(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+2:]...)
}

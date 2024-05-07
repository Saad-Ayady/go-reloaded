package mycode

import (
	"fmt"
	"strconv"
)

func Hex_Bin(s []string) []string {
	for i := 0 ; i < len(s) ; i++{
		tag := s[i]
		if tag == "(hex)"  && i-1 >= 0 && i < len(s){
			dec, err := strconv.ParseInt(s[i-1], 16, 64)
			
			if err != nil {
				fmt.Println("fuck error hh :)")
			}
			s[i-1] = strconv.Itoa(int(dec))
			s = rIndex(s, i)
		} else if tag == "(bin)" && i-1 >= 0 && i < len(s){
			dec, err := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = strconv.Itoa(int(dec))
			s = rIndex(s, i)
			if err != nil {
				fmt.Println("fuck error hh :)")
			}
		} else if tag == "(hex) " && i > 0 || tag == "(bin) " && i > 0{
			s = rIndex(s, i)
		}
	}
	return s
}

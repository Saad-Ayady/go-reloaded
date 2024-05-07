package mycode

import "strings"

func Apotrof(s []string) []string {
    apostropheCount := 0

    for i := 0; i < len(s); i++ {
        current := s[i]

        if strings.TrimSpace(current) == "'" { 
            if i+1 < len(s) && apostropheCount == 0 { 
                s[i+1] = "'" + strings.TrimSpace(s[i+1])
                s = append(s[:i], s[i+1:]...)            
                apostropheCount++                        
            } else if i-1 >= 0 { 
                s[i-1] = strings.TrimSpace(s[i-1]) + "'"
                apostropheCount = 0                      
                s = append(s[:i], s[i+1:]...)            
                i--                                      
            }
        }
    }
    return s
}
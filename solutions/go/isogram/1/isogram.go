package isogram

import (
    "slices"
    "strings"
) 


func IsIsogram(word string) bool {
    uniqueLetters := []rune{}

    lowercaseWord := strings.ToLower(word)

    for _, letter := range lowercaseWord {
        if slices.Contains(uniqueLetters, letter) {
            return false
        }
        if letter == '-' || letter == ' '{
            continue
        }
        uniqueLetters = append(uniqueLetters, letter)
    }
    
	return true 
}

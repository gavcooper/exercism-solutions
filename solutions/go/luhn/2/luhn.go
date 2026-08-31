package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	trimmedID := strings.ReplaceAll(id, " ", "")
	if len(trimmedID) <= 1 {
		return false
	}
	return luhnCheck(trimmedID)
}

func luhnCheck(id string) bool {
	var checkSum int

	needToDouble := len(id)%2 == 0 // an id of even length will have digits at even indices doubled, similarly odd length --> odd indices

	for _, digit := range id {
		if !unicode.IsDigit(digit) {
			return false
		}
		digit, _ := strconv.Atoi(string(digit))
		if needToDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		checkSum += digit
		needToDouble = !needToDouble
	}

	return checkSum%10 == 0
}

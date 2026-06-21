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

	if !containsDigitsOnly(trimmedID) {
		return false
	}

	return luhnCheck(trimmedID)
}

func containsDigitsOnly(id string) bool {
	for i := range id {
		character := rune(id[i])
		if !unicode.IsDigit(character) {
			return false
		}
	}
	return true
}

func luhnCheck(id string) bool {
	var checkDigits []int
	digits := strings.Split(id, "")

	var needToDouble bool

	for i := len(digits) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(digits[i])
		if needToDouble {
			digit = 2 * digit
			if digit > 9 {
				digit = digit - 9
			}
			checkDigits = append(checkDigits, digit)
			needToDouble = !needToDouble
			continue
		}
		checkDigits = append(checkDigits, digit)
		needToDouble = !needToDouble
	}

	var checkSum int

	for _, val := range checkDigits {
		checkSum += val
	}

	return checkSum%10 == 0

}

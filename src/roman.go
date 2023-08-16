package main

import (
	"fmt"
	"regexp"
)

// RomanToArabic converts a valid Roman numeral to an Arabic numeral
func RomanToArabic(roman string) string {
	if validateRomanNumber(roman) == 1 {
		return convert(roman)
	}
	return getMessage(INVALID_ROMAN_STRING)
}

func getValueFromRomanChar(romanChar byte) int {
	switch romanChar {
	case 'I':
		return int(I)
	case 'V':
		return int(V)
	case 'X':
		return int(X)
	case 'L':
		return int(L)
	case 'C':
		return int(C)
	case 'D':
		return int(D)
	case 'M':
		return int(M)
	}
	return -1
}

func validateRomanNumber(roman string) int {
	romanNumberValidator := "^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"
	if match, _ := regexp.MatchString(romanNumberValidator, roman); match {
		return 1
	}
	return 0
}

func convert(roman string) string {
	decimal := 0
	lastNumber := 0

	for i := len(roman) - 1; i >= 0; i-- {
		ch := roman[i]
		decimal = checkRoman(getValueFromRomanChar(ch), lastNumber, decimal)
		lastNumber = getValueFromRomanChar(ch)
	}

	return fmt.Sprintf("%d", decimal)
}

func checkRoman(totalDecimal, lastRomanLetter, lastDecimal int) int {
	if lastRomanLetter > totalDecimal {
		return lastDecimal - totalDecimal
	}
	return lastDecimal + totalDecimal
}
package propertybasedtests

import "strings"

// Roman numerals test: convert any arabic number to roman counterpart (e.g 4 => "IV")
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 9:
			result.WriteString("X")
			arabic -= 10
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	// TODO continue here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals#refactor-5

	return result.String()
}

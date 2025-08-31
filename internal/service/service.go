package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertText(input string) (string, error) {
	input = strings.TrimSpace(input)
	
	if isMorseCode(input) {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}

func isMorseCode(input string) bool {
	morseChars := ".-"
	spaceChars := " "
	
	for _, char := range input {
		if !strings.ContainsRune(morseChars+spaceChars, char) {
			return false
		}
	}
	
	return strings.Contains(input, ".") || strings.Contains(input, "-")
}

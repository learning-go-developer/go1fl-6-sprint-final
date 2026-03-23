package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectedMorse automatically detects whether the input string is Morse code or plain text
// and converts it to the opposite format.
//
// If the input string consists only of ".", "-", and spaces, it is treated as Morse code
// and converted into plain text. Otherwise, it is treated as plain text and
// encoded into Morse code.
//
// It returns the converted string and an error if the input is empty or contains only whitespace.
func DetectedMorse(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
	}

	if len(strings.TrimLeft(input, ".- ")) == 0 {
		return morse.ToText(input), nil // get text from morse
	}

	return morse.ToMorse(input), nil // get morse from text
}

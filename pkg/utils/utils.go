package utils

import (
	"strconv"
	"strings"
)

// Trim removes trailing spaces and symbols from the string.
func Trim(lexeme string) *string {
	str := lexeme
	str = strings.TrimFunc(str, func(r rune) bool {
		return r == '-' || r == '%' || r == '*' || r == '.' || r == ' '
	})

	if str == "" {
		return nil
	}

	return &str
}

// Atoi converts string into integer.
// Deprecated.
func Atoi(lexeme *string) (*int, error) {
	if lexeme == nil {
		return nil, nil
	}

	str := *lexeme
	str = strings.TrimFunc(str, func(r rune) bool {
		return r == '%' || r == '*' || r == '.' || r == ' '
	})

	if str == "" {
		return nil, nil
	}

	res, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Atof converts string into float.
// Deprecated.
func Atof(lexeme *string) (*float64, error) {
	if lexeme == nil {
		return nil, nil
	}

	str := *lexeme
	str = strings.TrimFunc(str, func(r rune) bool {
		return r == '%' || r == '*' || r == '.' || r == ' '
	})

	if str == "" {
		return nil, nil
	}

	str = strings.Replace(str, ",", ".", 1)
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

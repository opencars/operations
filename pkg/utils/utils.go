package utils

import (
	"strconv"
	"strings"
)

// Trim removes trailing spaces and symbols from the string.
func Trim(lexeme *string) *string {
	if lexeme == nil {
		return nil
	}

	str := *lexeme
	str = strings.TrimFunc(str, func(r rune) bool {
		return r == '-' || r == '%' || r == '*' || r == '.' || r == ' '
	})

	if str == "" {
		return nil
	}

	return &str
}

// Atoi converts string into integer.
func Atoi(lexeme *string) (*int, error) {
	if lexeme == nil {
		return nil, nil
	}

	str := *lexeme
	str = strings.TrimFunc(str, func(r rune) bool {
		return r == '%' || r == '*' || r == '.' || r == ' '
	})

	if str == "" || str == "NULL" {
		return nil, nil
	}

	res, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Atof converts string into float.
func Atof(lexeme *string) (*float64, error) {
	if lexeme == nil {
		return nil, nil
	}

	str := *lexeme
	str = strings.TrimFunc(str, func(r rune) bool {
		return r == '%' || r == '*' || r == '.' || r == ' '
	})

	if str == "" || str == "NULL" {
		return nil, nil
	}

	str = strings.Replace(str, ",", ".", 1)
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

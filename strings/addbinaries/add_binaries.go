package addbinaries

import (
	"errors"
)

const (
	empty = ""
)

var (
	errRuneNotInt = errors.New("digit is not an integer")
)

// Add adds two binary string numbers
func Add(number1 string, number2 string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getNumber(numberRunes []rune, index int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func characterToNumber(r rune) (int, error) { _ = "STUB: not implemented"; return 0, nil }

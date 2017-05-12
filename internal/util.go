package internal

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

func ValidName(nm string) error {
	var (
		index, count int
	)
	if len(nm) == 0 {
		return errors.New("names cannot be empty")
	}
	for index < len(nm) {
		rn, sz := utf8.DecodeRuneInString(nm[index:])

		if !unicode.IsLetter(rn) {
			if count < 1 {
				return errors.New("names must start with a letter")
			} else if rn != '_' || !unicode.IsNumber(rn) {
				return errors.New("names must only contain the underscore character, numbers, or letters")
			}
		}
		index += sz
		count++
	}
	return nil
}

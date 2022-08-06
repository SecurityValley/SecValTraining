package utils

import (
	"regexp"
	"strings"
)

func Clean(input string) string {
	pattern := regexp.MustCompile(`\s+`)

	return strings.ToLower(pattern.ReplaceAllString(input, " "))
}

func GetDefaultAlphabet() []rune {
	return []rune{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
	}
}

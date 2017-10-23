// Package strrev exports functions for reversing strings.
package strrev // import "4d63.com/strrev"

import (
	"unicode/utf8"
)

// Reverse reverses the given string, maintaining the bytes within any multibyte unicode characters in their existing order so that the character is still rendered correctly.
func Reverse(s string) string {
	reversed := make([]byte, len(s))

	for i := 0; i < len(reversed); {
		r, size := utf8.DecodeLastRuneInString(s)
		s = s[:len(s)-size]
		i += utf8.EncodeRune(reversed[i:], r)
	}

	return string(reversed)
}

// ReverseBytes reverses the given byte slice, maintaining the bytes within any multibyte unicode characters in their existing order so that the character is still rendered correctly.
func ReverseBytes(b []byte) []byte {
	reversed := make([]byte, len(b))

	for i := 0; i < len(reversed); {
		r, size := utf8.DecodeLastRune(b)
		b = b[:len(b)-size]
		i += utf8.EncodeRune(reversed[i:], r)
	}

	return reversed
}

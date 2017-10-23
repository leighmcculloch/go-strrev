// Strrev exports functions for reversing strings.
package strrev // import "4d63.com/strrev"

// Reverse reverses the given string, rune-by-rune, maintaining the bytes within any multibyte unicode characters in their existing order so that the character is still rendered correctly.
func Reverse(s string) string {
	r := []rune(s)
	length := len(r)
	last := length - 1
	for i := 0; i < length/2; i++ {
		r[i], r[last-i] = r[last-i], r[i]
	}
	return string(r)
}

// ReverseBytes reverses the given string, byte-by-byte. Multibyte unicode characters will have their internal bytes reversed also. In most cases you should use `Reverse` which maintains the bytes within a multibyte unicode character in the existing order so that the character is still rendered correctly.
func ReverseBytes(s string) string {
	b := []byte(s)
	length := len(b)
	last := length - 1
	for i := 0; i < length/2; i++ {
		b[i], b[last-i] = b[last-i], b[i]
	}
	return string(b)
}

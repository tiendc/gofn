//go:build go1.20

package gofn

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	omitAccentsTransformer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
)

// StringOmitAccents strips all accents from characters of a string
func StringOmitAccents(s string) (string, error) {
	for _, ch := range s {
		// If there's a character outside the range of ascii runes, there might be accented words
		if ch > 127 { // max standard ascii code
			ss, _, err := transform.String(omitAccentsTransformer, s)
			if err != nil {
				return "", err
			}
			return ss, nil
		}
	}
	return s, nil
}

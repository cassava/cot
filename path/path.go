package path

import (
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func Split(path string) []string {
	return strings.FieldsFunc(path, func(r rune) bool {
		return r == filepath.Separator
	})
}

// NOTE: If TranslatePrefix is handed an absolute string,
//		 you will NOT get a good result!
func TranslatePrefix(relpath string, prefix rune) string {
	xs := Split(relpath)
	for i, x := range xs {
		if len(x) <= 1 {
			// The string needs to consist of more than one character,
			// as we disallow translations of a string consisting only
			// of the prefix.
			continue
		}

		// Replace the prefix if necessary
		if r, w := utf8.DecodeRuneInString(x); r == prefix {
			xs[i] = "." + x[w:]
		}
	}
	return filepath.Join(xs...)
}

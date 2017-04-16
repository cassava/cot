package pathutil

import (
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// Split splits a string by path separator and returns a splice
// of path parts.
func Split(path string) []string {
	return strings.FieldsFunc(path, func(r rune) bool {
		return r == filepath.Separator
	})
}

// Translate prefix takes a prefix and replaces all instances
// of the prefix in a file with a '.'.
//
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

// FindRootDir returns the path of the first directory in the directory hierarchy
// where isRootFn returns true for that path.
//
// If isRootFn does not return true for any of the paths up to the global root "/",
// then the empty string is returned.
// If any error is returned by isRootFn, then that error is returned along with an
// empty string.
//
// TODO: stop at mount point like git does.
func FindRootDir(path string, isRootFn func(path string) (bool, error)) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	for {
		found, err := isRootFn(path)
		if err != nil {
			return "", err
		} else if found {
			return path, nil
		}

		// After we have checked / we stop looking.
		if path == "/" {
			break
		}
		path = filepath.Dir(path)
	}

	// Did not find the root.
	return "", nil
}

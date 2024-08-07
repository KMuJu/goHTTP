package goHTTP

import (
	"unicode"
)

var (
	Methods = []method{
		"GET", "POST", "HEAD", "PUT", "DELETE",
	}
)

func getMethod(header []byte) (method, int, error) {
	parsedMethod, n := getToWhiteSpace(header)
	for _, m := range Methods {
		if method(parsedMethod) == m {
			return m, n, nil
		}
	}

	return "", 0, methodNotSupported{}
}

func getToWhiteSpace(byteslice []byte) ([]byte, int) {
	for i, b := range byteslice {
		if unicode.IsSpace(rune(b)) {
			return byteslice[:i], i
		}
	}
	return byteslice, len(byteslice)
}

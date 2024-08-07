package internal

import (
	"slices"
	"strings"
	"unicode"
)

// Takes the input bytes and returns everything relevant on the request line.
// Request line is the first line of the request
//
// Ex:
// GET / HTTP/1.1
func parseRequestLine(input []byte) (length int, method string, url string, version string) {
	// TODO: Redo
	inputlen := len(input)
	start := length
	for ; length < inputlen; length++ { // Move length from start to first whitespace
		if unicode.IsSpace(rune(input[length])) {
			method = string(input[start:length]) // Method is then the string in [start, length)
			break
		}
	}
	length++ // Move past whitespace
	start = length
	for ; length < inputlen; length++ {
		if unicode.IsSpace(rune(input[length])) {
			url = string(input[start:length])
			break
		}
	}
	length++
	start = length
	for ; length < inputlen; length++ {
		if unicode.IsSpace(rune(input[length])) {
			version = string(input[start:length])
			break
		}
	}
	if length == inputlen && version == "" { // If the loop goes to the end, version needs to be set
		version = string(input[start:inputlen])
	}

	return
}

func parseHeaders(input []string) (Header, error) {
	headers := make(map[string][]string)
	for _, line := range input {
		h := strings.Split(line, ":")
		if len(h) < 2 { // header has format | Header: value1, value2
			return nil, errorInHeader
		}
		header := h[0]
		values := strings.Split(strings.Join(h[1:], ":"), ",")
		for i := 0; i < len(values); i++ {
			values[i] = strings.TrimSpace(values[i])
		}
		headers[header] = values
	}

	return headers, nil
}

func parseRequest(input []byte) (Request, error) {
	linelen, method, url, version := parseRequestLine(input)

	headerEnd := findHeaderEnd(input)

	// linelen + 2 to remove \r\n
	headerlines := strings.Split(string(input[linelen+2:headerEnd]), "\n")

	header, err := parseHeaders(headerlines)
	if err != nil {
		return Request{}, err
	}
	body := []byte{}
	if headerEnd+4 < len(input) { // +4 to hop over \r\n\r\n
		body = input[headerEnd+4:]
	}

	return Request{
		Method:  method,
		URL:     url,
		Version: version,
		Header:  header,
		Body:    body,
	}, nil
}

// Header ends with \r\n\r\n
func findHeaderEnd(input []byte) int {
	inputlen := len(input)
	for i, b := range input {
		if b == '\r' {
			if i+4 <= inputlen && slices.Equal(input[i:i+4], []byte{'\r', '\n', '\r', '\n'}) {
				return i
			}
		}
	}
	return -1
}

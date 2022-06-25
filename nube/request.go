package nube

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"golang.org/x/exp/slices"
)

var (
	REQUEST_METHODS = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
)

type Request struct {
	Method  string
	Host    string
	Path    string
	Header  Header
	Version string

	scanner *bufio.Scanner
}

func NewRequest(rd io.Reader) (*Request, error) {
	r := new(Request)
	r.scanner = bufio.NewScanner(rd)

	// Read request-line
	ok := r.scanner.Scan()

	if !ok {
		return nil, &RequestError{
			Message:    "Empty request",
			StatusCode: 400,
		}
	}

	var err error

	r.Method, r.Path, r.Version, err = parseRequestLine(r.scanner.Text())

	if err != nil {
		return nil, err
	}

	header, err := HeadersFromScanner(*r.scanner)
	
	if err != nil {
		return nil, err
	}

	r.Header = *header

	return r, nil
}

func parseRequestLine(line string) (string, string, string, error) {
	parts := strings.Split(line, SP)

	if len(parts) != 3 {
		return "", "", "", &RequestError{
			Message:    "Request line format is not okay, " + line,
			StatusCode: 400,
		}
	}

	method, path, version := parts[0], parts[1], parts[2]

	if slices.Index(REQUEST_METHODS, method) == -1 {
		return "", "", "", &RequestError{
			Message: "Invalid method " + method,
			StatusCode: 400,
		}
	}

	if !IsValidPath(path) {
		return "", "", "", &RequestError{
			Message: "Invalid path " + path,
			StatusCode: 400,
		}
	}

	return method, path, version, nil
}


func HeadersFromScanner(scanner bufio.Scanner) (*Header, error) {
	r := NewHeader()

	for scanner.Scan() {
		header := scanner.Text()

		if header == "" {
			break
		}

		parts := strings.Split(header, HEADER_SEP)
		key := parts[0]

		if !IsValidToken(key) {
			return nil, &RequestError{
				StatusCode: 400,
				Message: "Header field name is not a valid token: " + key,
			}
		}

		value := strings.Join(parts[1:], HEADER_SEP)

		r.headers[strings.ToLower(key)] = value
	}

	return r, nil
}

type RequestError struct {
	Message    string
	StatusCode int
}

func (err RequestError) Error() string {
	return fmt.Sprintf("status %d: %s", err.StatusCode, err.Message)
}

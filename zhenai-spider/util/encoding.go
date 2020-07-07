package util

import (
	"bufio"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}

	// DetermineEncoding determines the encoding of an HTML document by examining
	// up to the first 1024 bytes of content and the declared Content-Type.
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

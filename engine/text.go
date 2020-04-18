package engine

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// ScanHits is an alias for map[string]int, wich is the return type of a Scan
type ScanHits = map[string]int

// Scan : Counts number of each words written in a file
func Scan(r io.Reader) ScanHits {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	hits := make(map[string]int)
	for s.Scan() {
		hits[s.Text()]++
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return hits
}

// IsText : check that a file has a valid MIME text type
func IsText(r io.Reader) bool {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	t := http.DetectContentType(buf.Bytes())

	// TODO: change verification method
	return strings.HasPrefix(t, "text")
}
